package cmd

import (
	"fmt"

	"github.com/bekadoux/todo-cli/internal/todo"
	"github.com/spf13/cobra"
	//"text/tabwriter"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks",
	Long:  `List all tasks`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list called")
		fmt.Println(todo.LoadTasksFromCSV(todo.DefaultSavePath))
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
