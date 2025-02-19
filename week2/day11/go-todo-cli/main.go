package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

// Task represents a single to-do item
type Task struct {
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

// Filename for storing tasks
const taskFile = "tasks.json"

func LoadTasks() ([]Task, error) {
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

// SaveTasks writes tasks to a JSON file
func SaveTasks(tasks []Task) error {
	file, err := os.Create(taskFile)
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewEncoder(file).Encode(tasks)
}

// AddTask adds a new task to the list
func AddTask(description string) error {
	tasks, err := LoadTasks()
	if err != nil {
		return err
	}

	tasks = append(tasks, Task{Description: description, Completed: false})
	return SaveTasks((tasks))
}

// ListTasks displays all tasks
func ListTasks() {
	tasks, err := LoadTasks()
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}

	if len(tasks) == 0 {
		fmt.Println("No tasks found.")
		return
	}

	for i, task := range tasks {
		status := " "
		if task.Completed {
			status = "v"
		}
		fmt.Printf("%d. [%s] %s\n", i+1, status, task.Description)
	}
}

// MarkTaskAsDone marks a task as completed
func MarkTaskAsDone(index int) error {
	tasks, err := LoadTasks()
	if err != nil {
		return err
	}

	if index < 0 || index >= len(tasks) {
		return fmt.Errorf("Invalid task number")
	}

	tasks[index].Completed = true
	return SaveTasks(tasks)
}

// RemoveTask deletes a task from the list
func RemoveTask(index int) error {
	tasks, err := LoadTasks()
	if err != nil {
		return err
	}

	if index < 0 || index >= len(tasks) {
		return fmt.Errorf("Invalid task number")
	}

	tasks = append(tasks[:index], tasks[index+1:]...)
	return SaveTasks(tasks)
}

// CLI Menu
func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\nTo-Do List Manager")
		fmt.Println("1. Add Task")
		fmt.Println("2. List Tasks")
		fmt.Println("3. Mark Task as Completed")
		fmt.Println("4. Remove Task")
		fmt.Println("5. Exit")
		fmt.Print("Choose an option: ")

		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			fmt.Print("Enter task description: ")
			scanner.Scan()
			description := scanner.Text()
			err := AddTask(description)
			if err != nil {
				fmt.Println("Error adding task:", err)
			} else {
				fmt.Println("Task added successfully.")
			}

		case "2":
			ListTasks()

		case "3":
			fmt.Print("Enter task number to mark as completed: ")
			scanner.Scan()
			index, err := strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("Invalid input. Enter a valid task number.")
				continue
			}
			err = MarkTaskAsDone((index - 1))
			if err != nil {
				fmt.Println("Error marking task:", err)
			} else {
				fmt.Println("Task marked as completed.")
			}

		case "4":
			fmt.Print("Enter task number to remove: ")
			scanner.Scan()
			index, err := strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("Invalid input. Enter a valid task number.")
				continue
			}
			err = RemoveTask(index - 1)
			if err != nil {
				fmt.Println("Error removing task:", err)
			} else {
				fmt.Println("Task removed successfully!")
			}

		case "5":
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Please choose a valid option.")

		}
	}
}
