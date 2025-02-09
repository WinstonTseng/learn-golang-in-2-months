// This code demonstrates to append content instead of overwriting, use os.OpenFile() with os.O_APPEND.
// Date: February 9, 2025

package main

import (
	"fmt"
	"os"
)

func main() {
	// Open file in append mode
	file, err := os.OpenFile("output.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Failed to open the file:", err)
		return
	}

	defer file.Close()

	_, err = file.WriteString("This is appended content.\n")
	if err != nil {
		fmt.Println("Append failed:", err)
		return
	}

	fmt.Println("Append successful!")
}
