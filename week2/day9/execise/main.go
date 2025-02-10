package main

import (
	"execise/mathutils"
	"fmt"
)

func main() {
	product := mathutils.Multiply(10, 5)
	quotient, err := mathutils.Divide(10, 5)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Product of 10 * 5:", product)
		fmt.Println("Quotient of 10 / 5:", quotient)
	}

	product = mathutils.Multiply(10, 0)
	quotient, err = mathutils.Divide(10, 0)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Product of 10 * 0:", product)
		fmt.Println("Quotient of 10 / 0:", quotient)
	}

}
