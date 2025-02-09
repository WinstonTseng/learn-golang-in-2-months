// This code demonstrates how to count the number of lines in a file.
// Date: February 9, 2025

package main

import (
	"bufio"
	"fmt"
	"os"
)

func countLines(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineCount := 0
	for scanner.Scan() {
		lineCount++
	}

	if err = scanner.Err(); err != nil {
		return 0, err
	}

	return lineCount, nil
}

func main() {
	filename := "sample.txt"
	lines, err := countLines(filename)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("File %s has %d lines\n", filename, lines)
}
