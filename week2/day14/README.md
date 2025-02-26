## **Day 14: Review Week’s Progress and Refine CLI Project**

### **Objectives for Today:**
1. **Review and Reflect on Progress**  
   - Recap key concepts learned in the past week.
   - Identify strengths and areas needing improvement.
  
2. **Refine the CLI Project**  
   - Improve the command-line tool by optimizing code and adding features.
   - Implement error handling, logging, and code organization best practices.

3. **Write Unit Tests for the CLI Project**  
   - Ensure the application works as expected with automated tests.

---

### **1. Review of Key Concepts from Week 2**
This week, we focused on practical Go development and project structuring. Here’s a summary:

| Topic | Key Takeaways |
|---|---|
| **File Handling** | Reading and writing files using `os` and `bufio` packages. |
| **Packages and Modules** | `go mod init`, project structuring, and reusing code. |
| **Concurrency** | Introduction to Goroutines and Channels for parallel execution. |
| **CLI Project** | Built a to-do list application using Go. |
| **Testing & Debugging** | Writing unit tests with `testing` package, and debugging techniques. |

#### **Reflection Questions**
- What topics do you feel most confident about?
- What areas require more practice?
- Are there any gaps in understanding that need revision?

---

### **2. Refining the CLI Project**
Let's improve our CLI-based to-do list application by incorporating the following:
#### **(a) Enhancing User Input Handling**
Instead of relying on simple `fmt.Scanln()`, we can use `bufio.NewScanner` for better input processing:

```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getInput(prompt string) string {
	fmt.Print(prompt)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}
```

#### **(b) Error Handling**
Ensure we catch and handle errors properly:

```go
func loadTasks(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()
	
	var tasks []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		tasks = append(tasks, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}
	
	return tasks, nil
}
```

#### **(c) Implementing Logging**
Use the `log` package to track errors and actions:

```go
import "log"

func logError(err error) {
	if err != nil {
		log.Printf("Error: %v\n", err)
	}
}
```

---

### **3. Writing Unit Tests**
Since we created a CLI tool, we should test core functionalities like adding/removing tasks.

#### **(a) Writing a Test for Task Addition**
```go
package main

import (
	"testing"
)

func TestAddTask(t *testing.T) {
	tasks := []string{}
	tasks = addTask(tasks, "Buy groceries")

	if len(tasks) != 1 {
		t.Errorf("Expected 1 task, got %d", len(tasks))
	}

	if tasks[0] != "Buy groceries" {
		t.Errorf("Expected 'Buy groceries', got '%s'", tasks[0])
	}
}
```

#### **(b) Running Tests**
To run tests, execute:
```sh
go test ./...
```

---

### **4. Final Project Structure**
Refactor the project into multiple files:

```
cli-todo/
│── main.go        # Entry point
│── tasks.go       # Task management functions
│── utils.go       # Utility functions
│── tasks_test.go   # Unit tests for task functions
```

---

### **Wrap-up**
- Today, we reviewed past topics and refined our CLI project.
- We improved input handling, added logging, and wrote unit tests.
- If needed, you can extend the CLI tool by adding features like due dates, categories, or persistence using a database.

#### **Next Steps:**
Tomorrow, we move to **Phase 3**, where we start learning **HTTP servers and building APIs**! 🚀