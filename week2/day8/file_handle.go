// This file contains the code for file handling in Go.
// Golang uses os.Open() to open files and "bufio" or "ioutil" to read their contents.
// Date: February 9, 2025

// Explanation:

// os.Open("sample.txt") opens the file.
// bufio.NewScanner(file) reads the file line by line.
// scanner.Scan() moves to the next line, and scanner.Text() gets the line’s content.
// defer file.Close() ensures the file is closed after usage.

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// open the file
	file, err := os.Open("sample.txt")
	if err != nil {
		fmt.Println("Failed to open the file:", err)
		return
	}
	defer file.Close() // Ensure the file is colosed when the function exists

	// Read line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	// check for errors during reading
	if err := scanner.Err(); err != nil {
		fmt.Println("Error while reading file:", err)
	}
}
