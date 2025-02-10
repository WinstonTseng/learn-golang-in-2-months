package main

import (
	"fmt"
	"mymodule/mathops"

	"github.com/fatih/color"
)

func main() {
	sum := mathops.Add(10, 5)
	diff := mathops.Subtract(10, 5)

	fmt.Println("Sum:", sum)
	fmt.Println("Difference:", diff)

	color.Cyan("Hello, World in Cyan!")
	fmt.Println("This is normal text.")
}
