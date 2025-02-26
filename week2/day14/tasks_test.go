package main

import (
	"os"
	"testing"
)

func TestSaveAndLoadTasks(t *testing.T) {
	tmpFile := "test_tasks.txt"
	tasks := []string{"Task 1", "Task 2", "Task 3"}

	err := SaveTasks(tmpFile, tasks)
	if err != nil {
		t.Fatalf("Failed to save tasks: %v", err)
	}

	loadedTasks, err := LoadTasks(tmpFile)
	if err != nil {
		t.Fatalf("Failed to load tasks: %v", err)
	}

	if len(loadedTasks) != len(tasks) {
		t.Fatalf("Expected %d tasks, got %d", len(tasks), len(loadedTasks))
	}

	for i, task := range tasks {
		if loadedTasks[i] != task {
			t.Errorf("Expected task %q, got %q", task, loadedTasks[i])
		}
	}

	// Cleanup test file
	os.Remove(tmpFile)
}
