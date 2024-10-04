/*
Copyright © 2024 Harshwardhan
*/
package cmd

import (
	"fmt"

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
		userResp, err := datatypes.GetProfileResponse(username)
		if err != nil {
			fmt.Println(err)
		}
		repoResponse, err := datatypes.GetRepoResponse(username)

		user := datatypes.NewUser()
		user.SetUser(*userResp, *repoResponse)

		user.PrintWithStyle()
	},
}

func init() {
	rootCmd.AddCommand(profileCmd)
}
