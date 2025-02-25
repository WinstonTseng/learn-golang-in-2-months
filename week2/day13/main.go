package main

import (
	"day13/mathutil"
	"fmt"
)

func main() {
	a, b := 10, 5

	//log.Println("This is a log message")

	sum := mathutil.Add(a, b)
	sub := mathutil.Substract(a, b)

	fmt.Printf("Add(%d, %d) = %d\n", a, b, sum)
	fmt.Printf("Substract(%d, %d) = %d\n", a, b, sub)

	//log.Println("End of the program")
}
