/*
Copyright © 2024 Harshwardhan
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "GitTrack",
	Short: "Fetch Github user activity in terminal",
	Long:  ` A simple command-line tool that tracks and displays the latest GitHub activity of any user.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
