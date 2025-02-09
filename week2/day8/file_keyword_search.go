// Purpose: Go program to search for a keyword in a file.
// Date: February 9, 2025

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func searchKeyword(filename, keyword string) (bool, error) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Open file failed:", err)
		return false, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	totalCount := 0
	lineCount := 1
	for scanner.Scan() {
		line := scanner.Text()
		//fmt.Printf("Line %d: %s\n", lineCount, line) // Debug: Print each line
		if count := strings.Count(line, keyword); count > 0 {
			totalCount += count
			fmt.Printf("%s found in line %d\n", keyword, lineCount)
		}
		lineCount++
	}
	fmt.Printf("Total occurrences of %s: %d\n", keyword, totalCount)

	return true, nil
}

func main() {
	filename := "sample.txt"
	fmt.Println("Enter the keyword to search: ")
	var keyword string
	fmt.Scanln(&keyword)
	//fmt.Printf("Search for keyword: %s\n", keyword) // Debug: Print the keyword
	if _, err := searchKeyword(filename, keyword); err != nil {
		fmt.Println("Error:", err)
		return
	}
}
