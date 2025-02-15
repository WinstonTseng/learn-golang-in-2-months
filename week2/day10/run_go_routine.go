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
	go printMessage("Goroutine")  // Runs concurrently
	printMessage("Main Function") // Runs in the main thread
}
