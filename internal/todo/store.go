package todo

import (
	"encoding/csv"
	"fmt"
	"os"
)

var defaultSavePath string

func init() {
	var homeDir, err = os.UserHomeDir()
	if err != nil {
		panic("Failure getting user's home directory: " + err.Error())
	}
	defaultSavePath = fmt.Sprintf("%v/tasks.csv", homeDir)
}

func createOnNotExist(path string) error {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		file, err := os.Create(path)
		if err != nil {
			return fmt.Errorf("Error creating %s: %w", path, err)
		}
		defer file.Close()

		writer := csv.NewWriter(file)
		if err := writer.Write(getHeader()); err != nil {
			return fmt.Errorf("Error writing header: %w", err)
		}
		writer.Flush()
	}
	return nil
}

func SaveTaskToCSV(task Task) error {
	path := defaultSavePath
	err := createOnNotExist(path)
	if err != nil {
		fmt.Println("Failed to prepare file:", err)
		return err
	}
	file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Failed to open file:", err)
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	if err := writer.Write(task.getCSVData()); err != nil {
		fmt.Println("Failed writing to writer:", err)
		return err
	}
	writer.Flush()
	if err := writer.Error(); err != nil {
		fmt.Println("Failed writing to file: %w", err)
		return fmt.Errorf("Error writing task to file: %w", err)
	}
	return nil
}

//func loadTasksFromCSV() *[]Task {
//}
