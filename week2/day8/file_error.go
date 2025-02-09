// This code demonstrate how to handle file errors in Go.
// Date: February 9, 2025

package main

import (
	"fmt"
	"os"
)

func main() {
	_, err := os.Open("not_exist.txt")
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("File does not exist")
		} else if os.IsPermission(err) {
			fmt.Println("Permission denied")
		} else {
			fmt.Println("Error:", err)
		}

	}
}
