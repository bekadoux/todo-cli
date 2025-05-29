package todo

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

var DefaultSavePath string

func init() {
	var homeDir, err = os.UserHomeDir()
	if err != nil {
		panic("failure getting user's home directory: " + err.Error())
	}
	DefaultSavePath = fmt.Sprintf("%s/tasks.csv", homeDir)
}

func createOnNotExist(path string) error {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		file, err := os.Create(path)
		if err != nil {
			return fmt.Errorf("error creating %s: %w", path, err)
		}
		defer file.Close()

		writer := csv.NewWriter(file)
		if err := writer.Write(GetHeader()); err != nil {
			return fmt.Errorf("error writing header: %w", err)
		}
		writer.Flush()
	}
	return nil
}

func SaveNewTasksToCSV(manager *TaskManager) error {
	path := DefaultSavePath
	err := createOnNotExist(path)
	if err != nil {
		return fmt.Errorf("failed to prepare file: %w", err)
	}
	file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	if err := manager.ForEachNewTask(func(t *Task) error {
		if err := writer.Write(t.ToStringSlice()); err != nil {
			return fmt.Errorf("failed writing to writer: %w", err)
		}
		return nil
	}); err != nil {
		return fmt.Errorf("error saving task: %w", err)
	}
	writer.Flush()
	if err := writer.Error(); err != nil {
		return fmt.Errorf("error writing task to file: %w", err)
	}
	return nil
}

func LoadTasksFromCSV(manager *TaskManager, path string) error {
	file, err := os.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			return os.ErrNotExist // special case a user-friendly message
		}
		return fmt.Errorf("error opening %s: %w", path, err)
	}
	defer file.Close()
	reader := csv.NewReader(file)

	// Skip header
	_, err = reader.Read()
	if err != nil {
		return fmt.Errorf("error reading CSV from %s: %w", path, err)
	}
	csvContent, err := reader.ReadAll()
	if err != nil {
		return fmt.Errorf("error reading CSV from %s: %w", path, err)
	}

	for i, line := range csvContent {
		if len(line) < CSVTaskColumnCount {
			return fmt.Errorf("malformed line %d: expected %d columns, got %d", i+2, CSVTaskColumnCount, len(line))
		}

		id, err := strconv.Atoi(line[0])
		if err != nil {
			return fmt.Errorf("invalid ID at line %d: %w", i+2, err)
		}

		description := line[1]

		done, err := strconv.ParseBool(line[2])
		if err != nil {
			return fmt.Errorf("invalid Done value at line %d: %w", i+2, err)
		}

		manager.AddTaskFromStorage(id, description, done)
	}

	return nil
}
