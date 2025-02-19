## **Day 12: CLI Project Development (Part 2) – Enhancing the To-Do List Tool**

### **Goal**
- Continue building the **command-line To-Do List application**.
- Implement features for listing, removing, and marking tasks as completed.
- Persist tasks using a **JSON file**.
- Improve **user experience** with better input handling.

---

## **1. Review of Yesterday’s Progress**
Yesterday, you created the basic structure of the CLI tool:
- Defined a `Task` struct.
- Allowed users to add tasks.
- Implemented a basic CLI interface.

Today, you will:
1. Implement **task listing**.
2. Implement **task completion and removal**.
3. Add **file storage** for persistence.

---

## **2. Improving the To-Do List Tool**
Let’s extend the CLI application with more functionality.

### **2.1 Struct Definition**
Ensure the `Task` struct has fields for tracking completion and an ID.

```go
package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Task struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}
```

---

### **2.2 File Handling Functions**
Since we need to **persist tasks**, let's store them in a JSON file (`tasks.json`).

#### **2.2.1 Load Tasks from File**
```go
func loadTasks() ([]Task, error) {
	file, err := os.Open("tasks.json")
	if err != nil {
		if os.IsNotExist(err) {
			return []Task{}, nil // Return empty if file doesn't exist
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
```

#### **2.2.2 Save Tasks to File**
```go
func saveTasks(tasks []Task) error {
	file, err := os.Create("tasks.json")
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	return encoder.Encode(tasks)
}
```

---

### **2.3 Add a New Task**
Modify the **add task function** to store data persistently.
```go
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
```

---

### **2.4 List All Tasks**
```go
func listTasks() error {
	tasks, err := loadTasks()
	if err != nil {
		return err
	}

	if len(tasks) == 0 {
		fmt.Println("No tasks found.")
		return nil
	}

	for _, task := range tasks {
		status := "❌"
		if task.Completed {
			status = "✅"
		}
		fmt.Printf("[%d] %s %s\n", task.ID, status, task.Title)
	}
	return nil
}
```

---

### **2.5 Mark a Task as Completed**
```go
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
	return fmt.Errorf("task not found")
}
```

---

### **2.6 Remove a Task**
```go
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
		return fmt.Errorf("task not found")
	}

	return saveTasks(newTasks)
}
```

---

## **3. CLI Commands Handling**
Now, modify `main()` to allow CLI commands:

```go
package main

import (
	"fmt"
	"os"
	"strconv"
)

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
```

---

## **4. Testing Your CLI Tool**
### **Adding Tasks**
```sh
go run main.go add "Buy groceries"
go run main.go add "Read a book"
```

### **Listing Tasks**
```sh
go run main.go list
```
Example output:
```
[1] ❌ Buy groceries
[2] ❌ Read a book
```

### **Marking a Task as Completed**
```sh
go run main.go done 1
```

### **Removing a Task**
```sh
go run main.go remove 2
```

---

## **5. Summary of Today's Learning**
✅ Implemented **file persistence** using JSON.  
✅ Added **list, complete, and remove** features.  
✅ Improved **error handling** and **user experience**.  

---

## **6. Next Steps**
Tomorrow, we will:
- **Write unit tests** for this CLI tool.
- Improve error handling further.
- Package the tool for easy use.

