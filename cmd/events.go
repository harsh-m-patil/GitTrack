/*
Copyright © 2024 Harshwardhan
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// eventsCmd represents the events command
var eventsCmd = &cobra.Command{
	Use:   "events",
	Short: "Show events of the user",
	Long:  `Show events of the user`,

	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(args[0])
		fmt.Println("events called")
	},
}

func init() {
	rootCmd.AddCommand(eventsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// eventsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// eventsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
