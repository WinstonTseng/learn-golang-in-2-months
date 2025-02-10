### **Day 9: Packages and Modules in Go**
#### **Goal:**  
- Understand how to create and use packages in Go.  
- Learn how to structure Go projects using modules (`go mod`).  
- Practice by creating a simple package for basic mathematical operations.  

---

### **1. Introduction to Go Modules**
Go uses modules to manage dependencies and package organization. A **module** is a collection of Go packages, and it is defined by a `go.mod` file.

#### **Why Use Modules?**
- Helps organize code into reusable components.
- Simplifies dependency management.
- Provides version control over external libraries.

#### **Initializing a Module**
To create a new module, use:

```sh
go mod init mymodule
```
This creates a `go.mod` file with the module name.

---

### **2. Creating and Using Packages**
A **package** is a directory with one or more `.go` files. The file must start with a `package` declaration.

#### **Steps to Create a Package:**
1. Create a new directory for the project.
2. Inside the directory, create a subdirectory (e.g., `mathops/`) to hold the package.
3. Inside `mathops/`, create a Go file (e.g., `math.go`) and define the package.

#### **Example: Creating a Math Package**
**Folder structure:**
```
mymodule/
│── go.mod
│── main.go
│── mathops/
│   ├── math.go
```

**mathops/math.go**
```go
package mathops

// Add function adds two numbers
func Add(a, b int) int {
    return a + b
}

// Subtract function subtracts two numbers
func Subtract(a, b int) int {
    return a - b
}
```

**main.go**
```go
package main

import (
    "fmt"
    "mymodule/mathops"
)

func main() {
    sum := mathops.Add(10, 5)
    diff := mathops.Subtract(10, 5)

    fmt.Println("Sum:", sum)
    fmt.Println("Difference:", diff)
}
```

#### **Running the Code**
```sh
go run main.go
```
**Expected Output:**
```
Sum: 15
Difference: 5
```

---

### **3. Importing External Packages**
Go allows importing external packages using `go get`.

#### **Example: Using an External Library**
To install a library (e.g., `github.com/fatih/color` for colored output):

```sh
go get github.com/fatih/color
```

Then, use it in your code:

```go
package main

import (
    "fmt"
    "github.com/fatih/color"
)

func main() {
    color.Cyan("Hello, World in Cyan!")
    fmt.Println("This is normal text.")
}
```

Run the program:
```sh
go run main.go
```

---

### **4. Organizing Large Projects**
For large projects, follow these conventions:
- **cmd/**: Contains main entry points for CLI applications.
- **internal/**: Contains private packages for internal use.
- **pkg/**: Contains reusable packages.
- **api/**: Stores API-related code (for web apps).
- **configs/**: Stores configuration files.

Example:
```
myproject/
│── go.mod
│── cmd/
│   ├── app.go
│── internal/
│   ├── database/
│── pkg/
│   ├── mathops/
│── api/
│   ├── routes/
│── configs/
│   ├── config.yaml
```

---

### **5. Exercise: Create a Custom Math Package**
#### **Task:**  
1. Create a module `mathutils`.
2. Inside `mathutils`, create a package `operations` with functions:
   - `Multiply(a, b int) int`
   - `Divide(a, b int) (int, error)` (Handle division by zero).
3. Use the package in `main.go` and test it.

---
