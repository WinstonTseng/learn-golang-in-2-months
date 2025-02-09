package main

import (
	"fmt"
	"math"
)

func IsPrime(n int) bool {
	if n <= 2 {
		return false
	}

	if n%2 == 0 {
		return false
	}

	limit := int(math.Sqrt(float64(n)))

	for i := 3; i <= limit; i += 2 {
		if n%i == 0 {
			return false
		}

	}
	return true
}

func main() {
	numbers := []int{1, 2, 3, 4, 11, 12, 13, 14, 21, 22, 23}

	for _, num := range numbers {
		fmt.Printf("%d is Prime numbers %v\n", num, IsPrime(num))
	}
}
