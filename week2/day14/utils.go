package main

import (
	"bufio"
	"fmt"
	"os"
)

// getInput reads user input from the console
func getInput(prompt string) string {
	fmt.Print(prompt)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

// logError logs an error message if an error occurs
func logError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
	}
}
