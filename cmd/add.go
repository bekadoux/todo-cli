package cmd

import (
	"encoding/csv"
	"fmt"
	"github.com/bekadoux/todo-cli/internal/todo"
	"github.com/spf13/cobra"
	"os"
	"strconv"
	"strings"
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

		saveToStore(newTask)
	},
}

func saveToStore(task todo.Task) {
	records := [][]string{{"ID", "Description", "Done"},
		{strconv.Itoa(int(task.ID)), task.Description, strconv.FormatBool(task.Done)}}
	w := csv.NewWriter(os.Stdout)

	for _, record := range records {
		if err := w.Write(record); err != nil {
			fmt.Println("error")
		}
	}

	w.Flush()
}

func init() {
	rootCmd.AddCommand(addCmd)
}
