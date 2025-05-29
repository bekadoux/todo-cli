package cmd

import (
	"fmt"

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
		manager := todo.NewTaskManager()

		loadPath := todo.DefaultSavePath
		if err := todo.LoadTasksFromCSV(manager, loadPath); err != nil {
			if err == os.ErrNotExist {
				fmt.Printf(`
It looks like there are no tasks yet, and the file %q does not exist.

- To add your first task, use: todo-cli add
- To see all available commands, run: todo-cli help
- If you specified a custom file path, double-check that the path is correct.

If you believe the file should exist but it's missing, ensure it wasn't accidentally deleted or moved.
`, loadPath)
				return
			}
			fmt.Printf("error loading tasks from CSV: %v\n", err)
			os.Exit(1)
		}

		fmt.Fprintln(w, strings.Join(todo.GetHeader(), "\t"))
		manager.ForEachTask(func(t *todo.Task) {
			fmt.Fprintln(w, strings.Join(t.ToStringSlice(), "\t"))
		})

		w.Flush()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
