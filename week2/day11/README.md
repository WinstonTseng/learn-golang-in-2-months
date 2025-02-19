## Day 11: Work on a CLI Project – Part 1

### **Goal:**  
Today, you'll start building a **command-line to-do list manager** in Golang. The project will involve reading user inputs, storing tasks, and managing a simple task list.

---

### **1. Features of the CLI To-Do List**
We'll implement the following features:
- **Add** a task.
- **List** all tasks.
- **Remove** a task by index.
- **Mark** a task as completed.

---

### **2. Setting Up the Project**
1. Create a new directory for the project:
   ```sh
   mkdir go-todo-cli && cd go-todo-cli
   ```
2. Initialize a Go module:
   ```sh
   go mod init go-todo-cli
   ```

---

### **3. Implementing the CLI Logic**
Create a file named **main.go** and add the following code:

```go
package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Task represents a single to-do item
type Task struct {
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

// Filename for storing tasks
const taskFile = "tasks.json"

// LoadTasks reads tasks from a JSON file
func LoadTasks() ([]Task, error) {
	file, err := os.Open(taskFile)
	if err != nil {
		if os.IsNotExist(err) {
			return []Task{}, nil // Return an empty list if the file doesn't exist
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
	return SaveTasks(tasks)
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
		status := "❌"
		if task.Completed {
			status = "✅"
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
		return fmt.Errorf("invalid task number")
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
		return fmt.Errorf("invalid task number")
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
				fmt.Println("Task added successfully!")
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
			err = MarkTaskAsDone(index - 1)
			if err != nil {
				fmt.Println("Error marking task:", err)
			} else {
				fmt.Println("Task marked as completed!")
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
			fmt.Println("Invalid choice. Please select a valid option.")
		}
	}
}
```

---

### **4. Running and Testing**
To run the CLI application:
```sh
go run main.go
```
You can interact with the menu:
1. Add a task (e.g., "Buy groceries").
2. List tasks.
3. Mark a task as completed.
4. Remove a task.
5. Exit the program.

---

### **5. Next Steps (Day 12)**
- Improve error handling.
- Add persistent storage using BoltDB.
- Implement command-line arguments instead of interactive input.
- Package the application for distribution.

---

