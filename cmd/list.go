package cmd

import (
	"fmt"

	//"github.com/bekadoux/todo-cli/internal/todo"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/bekadoux/todo-cli/internal/todo"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks",
	Long:  `List all tasks`,
	Run: func(cmd *cobra.Command, args []string) {
		w := tabwriter.NewWriter(os.Stdout, 10, 4, 10, ' ', 0)
		tasks, err := todo.LoadTasksFromCSV(todo.DefaultSavePath)
		if err != nil {
			fmt.Printf("error loading tasks from CSV: %v\n", err)
			os.Exit(1)
		}

		fmt.Fprintln(w, strings.Join(todo.GetHeader(), "\t"))
		for _, task := range tasks {
			fmt.Fprintln(w, strings.Join(task.ToStringSlice(), "\t"))
		}

		w.Flush()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
