package cmd

import (
	"fmt"
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
		description := strings.Join(args, " ")
		if len(description) == 0 {
			fmt.Printf("no task description provided\n")
			return
		}

		const maxDescriptionLength = 100
		if len(description) > maxDescriptionLength {
			fmt.Printf("task description too long (max %d characters, got %d)\n", maxDescriptionLength, len(args))
			return
		}

		var newTask todo.Task = todo.Task{
			ID:          0,
			Description: description,
			Done:        false,
		}

		todo.SaveTaskToCSV(newTask)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
