package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

// Task represents a single to-do item
type Task struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

// Filename for storing tasks
const taskFile = "tasks.json"

func loadTasks() ([]Task, error) {
	file, err := os.Open(taskFile)
	if err != nil {
		if os.IsNotExist(err) {
			return []Task{}, nil
		}
		return nil, err
	}
	defer file.Close()

	var tasks []Task
	err = json.NewDecoder(file).Decode(&tasks)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

// saveTasks writes tasks to a JSON file
func saveTasks(tasks []Task) error {
	file, err := os.Create(taskFile)
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewEncoder(file).Encode(tasks)
}

// addTask adds a new task to the list
func addTask(title string) error {
	tasks, err := loadTasks()
	if err != nil {
		return err
	}

	// Assign an ID
	newTask := Task{
		ID:    len(tasks) + 1,
		Title: title,
	}

	tasks = append(tasks, newTask)
	return saveTasks(tasks)
}

// listTasks displays all tasks
func listTasks() error {
	tasks, err := loadTasks()
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return err
	}

	if len(tasks) == 0 {
		fmt.Println("No tasks found.")
		return nil
	}

	for _, task := range tasks {
		status := " "
		if task.Completed {
			status = "v"
		}
		fmt.Printf("%d. [%s] %s\n", task.ID, status, task.Title)
	}
	return nil
}

func completeTask(id int) error {
	tasks, err := loadTasks()
	if err != nil {
		return err
	}

	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Completed = true
			return saveTasks(tasks)
		}
	}
	return fmt.Errorf("Task not found")
}

// removeTask deletes a task from the list
func removeTask(id int) error {
	tasks, err := loadTasks()
	if err != nil {
		return err
	}

	newTasks := []Task{}
	for _, task := range tasks {
		if task.ID != id {
			newTasks = append(newTasks, task)
		}
	}

	if len(newTasks) == len(tasks) {
		return fmt.Errorf("Task not found")
	}

	return saveTasks(newTasks)
}

// CLI Menu

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: todo [add|list|done|remove] [task]")
		return
	}

	command := os.Args[1]

	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Error: Please provide a task description.")
			return
		}
		err := addTask(os.Args[2])
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("Task added successfully.")
		}

	case "list":
		err := listTasks()
		if err != nil {
			fmt.Println("Error:", err)
		}

	case "done":
		if len(os.Args) < 3 {
			fmt.Println("Error: Please provide a task ID.")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Error: Invalid task ID.")
			return
		}
		err = completeTask(id)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("Task marked as completed.")
		}

	case "remove":
		if len(os.Args) < 3 {
			fmt.Println("Error: Please provide a task ID.")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Error: Invalid task ID.")
			return
		}
		err = removeTask(id)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("Task removed successfully.")
		}

	default:
		fmt.Println("Unknown command. Available commands: add, list, done, remove.")
	}
}
