package main

func Fibonacci(i int) int {
	if i == 0 {
		return 0
	}
	if i == 1 {
		return 1
	}

	a, b := 0, 1

	for n := 2; n <= i; n++ {
		a, b = b, a+b
	}
	return b

}
