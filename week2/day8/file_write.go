// This code demonstrates that os.Create() creates or overwrites a file, while os.OpenFile() is used for appending content.
// Date: February 9, 2025

package main

import (
	"fmt"
	"os"
)

func main() {
	// Create and write to a file
	file, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Failed to create the file:", err)
		return
	}
	defer file.Close()

	// write contest
	_, err = file.WriteString("Hello, Golang!\n")
	if err != nil {
		fmt.Println("Write failed:", err)
		return
	}

	fmt.Println("Write successfule!")

}
