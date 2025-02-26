package main

import (
	"bufio"
	"os"
)

func LoadTasks(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return []string{}, nil
		}
		return nil, err
	}
	defer file.Close()

	var tasks []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		tasks = append(tasks, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

func SaveTasks(filename string, tasks []string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, task := range tasks {
		_, err := writer.WriteString(task + "\n")
		if err != nil {
			return err
		}
	}

	return writer.Flush()
}
