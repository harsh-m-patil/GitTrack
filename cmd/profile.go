/*
Copyright © 2024 Harshwardhan
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
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
		profileURL := fmt.Sprintf(datatypes.UserProfileURL, username)

		resp, err := http.Get(profileURL)
		if err != nil {
			fmt.Printf("Error Making GET request: %v\n", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode == http.StatusNotFound {
			fmt.Println("User not found")
		}

		if resp.StatusCode != http.StatusOK {
			fmt.Printf("Failed to get a valid response,Status : %d\n", resp.StatusCode)
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("Error reading response body %v\n", err)
		}

		var userResponse datatypes.UserResponse

		err = json.Unmarshal(body, &userResponse)
		if err != nil {
			log.Printf("Error Unmarshalling JSON %v\n", err)
		}

		printProfile(userResponse)
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

func printProfile(user datatypes.UserResponse) {
	fmt.Printf("UserName : %s\n", user.Login)
	fmt.Printf("Name : %s\n", user.Name)
	fmt.Printf("Public Repos : %d\n", user.PublicRepos)
	fmt.Printf("Followers : %d\n", user.Followers)
	fmt.Printf("Following : %d\n", user.Following)
}
