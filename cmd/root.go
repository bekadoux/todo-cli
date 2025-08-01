package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "todo-cli",
	Short: "A CLI To-Do list application written in Go",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		listCmd.Run(cmd, []string{})
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
