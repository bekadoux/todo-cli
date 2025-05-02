package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// doneCmd represents the done command
var doneCmd = &cobra.Command{
	Use:   "done",
	Short: "mark task as done",
	Long: `Mark task as completed. Usage:
	todo-cli done your_task`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("done called")
	},
}

func init() {
	rootCmd.AddCommand(doneCmd)
}
