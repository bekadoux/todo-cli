package todo

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	// "strings"
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
		if err := writer.Write(getHeader()); err != nil {
			return fmt.Errorf("error writing header: %w", err)
		}
		writer.Flush()
	}
	return nil
}

func SaveTaskToCSV(task Task) error {
	path := DefaultSavePath
	err := createOnNotExist(path)
	if err != nil {
		fmt.Println("failed to prepare file:", err)
		return err
	}
	file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("failed to open file:", err)
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	if err := writer.Write(task.getCSVData()); err != nil {
		fmt.Println("failed writing to writer:", err)
		return err
	}
	writer.Flush()
	if err := writer.Error(); err != nil {
		fmt.Println("failed writing to file: %w", err)
		return fmt.Errorf("error writing task to file: %w", err)
	}
	return nil
}

func LoadTasksFromCSV(path string) ([]Task, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("error opening %s: %w", path, err)
	}
	defer file.Close()
	reader := csv.NewReader(file)

	// Skip header
	_, err = reader.Read()
	if err != nil {
		return nil, fmt.Errorf("error reading CSV from %s: %w", path, err)
	}
	csvContent, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("error reading CSV from %s: %w", path, err)
	}

	var tasks []Task
	for i, line := range csvContent {
		const expectedColumnCount = 3
		if len(line) < expectedColumnCount {
			return nil, fmt.Errorf("malformed line %d: expected %d columns, got %d", i+2, expectedColumnCount, len(line))
		}

		id, err := strconv.Atoi(line[0])
		if err != nil {
			return nil, fmt.Errorf("invalid ID at line %d: %w", i+2, err)
		}

		description := line[1]

		done, err := strconv.ParseBool(line[2])
		if err != nil {
			return nil, fmt.Errorf("invalid Done value at line %d: %w", i+2, err)
		}

		tasks = append(tasks, Task{
			ID:          id,
			Description: description,
			Done:        done,
		})
	}

	return tasks, nil
}
