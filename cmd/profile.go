/*
Copyright © 2024 Harshwardhan
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// profileCmd represents the profile command
var profileCmd = &cobra.Command{
	Use:   "profile",
	Short: "Show github user's profile",
	Long:  `Display users github info`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(args[0])
		fmt.Println("profile called")
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
