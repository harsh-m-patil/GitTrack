/*
Copyright © 2024 Harshwardhan
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/harsh-m-patil/GitTrack/datatypes"
	"github.com/spf13/cobra"
)

// profileCmd represents the profile command
var profileCmd = &cobra.Command{
	Use:   "profile",
	Short: "Show github user's profile",
	Long:  `Display users github info`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		username := args[0]
		userResp, err := getProfileResponse(username)
		if err != nil {
			fmt.Println(err)
		}
		repoResponse, err := getRepoResponse(username)

		user := datatypes.NewUser()
		user.SetUser(*userResp, *repoResponse)

		fmt.Println(user.ToString())
	},
}

func init() {
	rootCmd.AddCommand(profileCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// profileCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// profileCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func getProfileResponse(username string) (*datatypes.UserResponse, error) {
	profileURL := fmt.Sprintf(datatypes.UserProfileURL, username)

	resp, err := http.Get(profileURL)
	if err != nil {
		return nil, fmt.Errorf("error making GET request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return nil, fmt.Errorf("User not found")
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Failed to get a valid response,Status : %d\n", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Error reading response body %v\n", err)
	}

	var userResponse datatypes.UserResponse

	err = json.Unmarshal(body, &userResponse)
	if err != nil {
		return nil, fmt.Errorf("Error Unmarshalling JSON %v\n", err)
	}

	return &userResponse, nil
}

func getRepoResponse(username string) (*datatypes.RepoResponse, error) {
	reposURL := fmt.Sprintf(datatypes.ReposURL, username)

	resp, err := http.Get(reposURL)
	if err != nil {
		return nil, fmt.Errorf("error making GET request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return nil, fmt.Errorf("User not found")
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Failed to get a valid response,Status : %d\n", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Error reading response body %v\n", err)
	}

	var repoResponse datatypes.RepoResponse

	err = json.Unmarshal(body, &repoResponse)
	if err != nil {
		return nil, fmt.Errorf("Error Unmarshalling JSON %v\n", err)
	}

	return &repoResponse, nil
}
