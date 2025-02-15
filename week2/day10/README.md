## **Day 10: Introduction to Concurrency in Golang**

### **1. Learning Objectives**
By the end of today, you should be able to:
- Understand what concurrency is and how it benefits Go applications.
- Use **goroutines** to execute functions concurrently.
- Use **channels** to communicate between goroutines.
- Implement **sync.WaitGroup** to manage multiple goroutines.

---

### **2. Introduction to Concurrency**
Golang has built-in support for concurrency using **goroutines** and **channels**. Concurrency allows multiple functions to run independently, making programs faster and more efficient.

#### **2.1 What is a Goroutine?**
A **goroutine** is a lightweight thread managed by the Go runtime. You can create one by using the `go` keyword before a function call.

#### **2.2 What are Channels?**
Channels are used to communicate between goroutines safely. They help prevent race conditions and ensure proper data synchronization.

---

### **3. Implementing Goroutines**
#### **Example 1: Running a Function as a Goroutine (run_go_routine.go) **
```go
package main

import (
	"fmt"
	"time"
)

func printMessage(msg string) {
	for i := 0; i < 5; i++ {
		fmt.Println(msg, i)
		time.Sleep(time.Millisecond * 500) // Simulate work
	}
}

func main() {
	go printMessage("Goroutine") // Runs concurrently
	printMessage("Main Function") // Runs in the main thread
}
```
**Output (may vary due to concurrency):**
```
Main Function 0
Goroutine 0
Main Function 1
Goroutine 1
...
```
> **Note:** The main function may exit before the goroutine completes execution.

---

### **4. Using WaitGroup to Synchronize Goroutines**
Since goroutines run asynchronously, we often need to **wait** for them to complete. The `sync.WaitGroup` package allows us to achieve this.

#### **Example 2: Using sync.WaitGroup**
```go
package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Notify when done
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	wg.Wait() // Wait for all goroutines to finish
	fmt.Println("All workers finished")
}
```
**Output:**
```
Worker 1 starting
Worker 2 starting
Worker 3 starting
Worker 1 done
Worker 2 done
Worker 3 done
All workers finished
```
> **Key Points:**
> - `wg.Add(1)`: Increments the counter before starting a goroutine.
> - `wg.Done()`: Decrements the counter when a goroutine completes.
> - `wg.Wait()`: Blocks execution until all goroutines finish.

---

### **5. Using Channels for Communication**
#### **Example 3: Sending Data Through a Channel**
```go
package main

import (
	"fmt"
)

func sendData(ch chan string) {
	ch <- "Hello from Goroutine"
}

func main() {
	messageChannel := make(chan string) // Create a channel

	go sendData(messageChannel) // Run in a goroutine

	message := <-messageChannel // Receive data from channel
	fmt.Println(message)
}
```
**Output:**
```
Hello from Goroutine
```
> **Key Points:**
> - `make(chan string)`: Creates a channel.
> - `ch <- "Hello"`: Sends data into the channel.
> - `<-ch`: Receives data from the channel.

---

### **6. Exercise: Parallel Calculation of Squares**
#### **Problem Statement**
Write a Go program that calculates the squares of numbers **concurrently** and sends the results to a channel.

#### **Solution**
```go
package main

import (
	"fmt"
	"sync"
)

func square(num int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	ch <- num * num
}

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	ch := make(chan int, len(numbers)) // Buffered channel
	var wg sync.WaitGroup

	for _, num := range numbers {
		wg.Add(1)
		go square(num, ch, &wg)
	}

	wg.Wait()
	close(ch)

	for result := range ch {
		fmt.Println(result)
	}
}
```
**Expected Output:**
```
4
16
36
64
100
```
> **Concepts Used:**
> - Using **goroutines** for parallel computation.
> - Using **channels** to collect results.
> - Using **sync.WaitGroup** to wait for all computations.

---

### **7. Summary**
1. **Goroutines** allow functions to execute concurrently using the `go` keyword.
2. **WaitGroup** ensures that the main function waits for goroutines to complete.
3. **Channels** facilitate safe communication between goroutines.
4. Combining **WaitGroup** and **channels** enables effective concurrency control.

---

### **8. Next Steps**
✅ **Practice:** Modify the square calculation program to also compute cube values in parallel.  
🚀 **Tomorrow (Day 11-12):** Start working on a **CLI project** using concurrency to manage tasks.

