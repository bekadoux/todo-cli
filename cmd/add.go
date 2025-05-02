package cmd

import (
	"strings"

	"github.com/spf13/cobra"

	"github.com/bekadoux/todo-cli/internal/todo"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task",
	Long: `Add a new task to the list. Usage:
	todo-cli add your-task`,
	Run: func(cmd *cobra.Command, args []string) {
		var newTask todo.Task = todo.Task{
			ID:          0,
			Description: strings.Join(args, " "),
			Done:        false,
		}

		todo.SaveTaskToCSV(newTask)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
