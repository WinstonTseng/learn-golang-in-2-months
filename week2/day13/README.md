### **Day 13: Testing and Debugging in Go**

Today, you’ll focus on writing unit tests in Go and debugging techniques. Testing ensures code reliability, while debugging helps identify and fix issues efficiently.

---

## **1. Importance of Testing**
Testing is crucial to ensure your Go programs work as expected. Go has a built-in testing framework under the `testing` package.

### **Types of Testing in Go**
1. **Unit Testing** – Testing small, individual components (functions, methods).
2. **Integration Testing** – Testing interactions between components.
3. **Benchmarking** – Measuring the performance of functions.

---

## **2. Writing Unit Tests in Go**
### **Basic Structure of a Test File**
- Test files should be named with `_test.go` (e.g., `math_test.go`).
- Each test function starts with `Test` and takes `*testing.T` as a parameter.

### **Example: Unit Test for a Math Function**
#### **math.go**
```go
package mathutil

// Add returns the sum of two numbers
func Add(a, b int) int {
    return a + b
}

// Subtract returns the difference of two numbers
func Subtract(a, b int) int {
    return a - b
}
```

#### **math_test.go**
```go
package mathutil

import "testing"

func TestAdd(t *testing.T) {
    result := Add(2, 3)
    expected := 5
    if result != expected {
        t.Errorf("Add(2, 3) = %d; want %d", result, expected)
    }
}

func TestSubtract(t *testing.T) {
    result := Subtract(10, 5)
    expected := 5
    if result != expected {
        t.Errorf("Subtract(10, 5) = %d; want %d", result, expected)
    }
}
```

### **Running Tests**
Run the tests using:
```sh
go test
```
or for verbose output:
```sh
go test -v
```

---

## **3. Table-Driven Testing**
Instead of writing multiple test cases manually, you can use **table-driven tests**.

```go
func TestAddTableDriven(t *testing.T) {
    tests := []struct {
        a, b     int
        expected int
    }{
        {2, 3, 5},
        {10, -2, 8},
        {0, 0, 0},
    }

    for _, test := range tests {
        result := Add(test.a, test.b)
        if result != test.expected {
            t.Errorf("Add(%d, %d) = %d; want %d", test.a, test.b, result, test.expected)
        }
    }
}
```

---

## **4. Writing Benchmark Tests**
Go’s testing framework includes support for benchmarking functions.

```go
func BenchmarkAdd(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Add(10, 20)
    }
}
```
Run benchmark tests:
```sh
go test -bench .
```

---

## **5. Debugging Techniques**
### **Using fmt.Println for Quick Debugging**
Adding `fmt.Println` statements is a simple but effective debugging technique.

```go
fmt.Println("Debug: Value of x =", x)
```

### **Using log Package**
`log` provides more control over logging messages.

```go
import "log"

log.Println("This is a log message")
log.Fatalf("Critical error occurred: %v", err)
```

### **Using Go’s Built-in Debugger (Delve)**
[Delve](https://github.com/go-delve/delve) is a powerful Go debugger.

#### **Install Delve**
```sh
go install github.com/go-delve/delve/cmd/dlv@latest
```

#### **Run a Go Program in Debug Mode**
```sh
dlv debug main.go
```

### **Setting Breakpoints and Inspecting Variables**
Inside `dlv`:
```sh
break main.go:10   # Set a breakpoint at line 10
continue           # Run until breakpoint
print x            # Print variable x
next               # Step to next line
```

---

## **6. Writing Tests for CLI Programs**
For CLI programs, use `os/exec` to simulate command-line execution.

```go
import (
    "os/exec"
    "testing"
)

func TestCLI(t *testing.T) {
    cmd := exec.Command("./mycli", "arg1")
    output, err := cmd.Output()
    if err != nil {
        t.Fatalf("Error running CLI command: %v", err)
    }

    expected := "Expected Output"
    if string(output) != expected {
        t.Errorf("CLI output = %s; want %s", output, expected)
    }
}
```

---

## **7. Debugging a CLI Program**
When debugging a CLI tool, you can run it interactively with:
```sh
dlv exec ./mycli
```

Or add logging to your program to print debug info:
```go
log.Printf("Processing input: %s", input)
```

---

## **8. Homework**
1. **Write unit tests for your CLI project.** 
   - Ensure key functions are tested.
   - Use table-driven tests.
2. **Run and analyze benchmarks.** 
   - Optimize your CLI tool’s performance if needed.
3. **Install and practice debugging with Delve.**
   - Set breakpoints and step through your code.

---

### **Next Step: Day 14 - Review and Refining the CLI Project**
Tomorrow, we will review your CLI project, improve its structure, and prepare it for final testing and optimization. 🚀