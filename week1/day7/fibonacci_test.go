package main

import "testing"

func TestFibonacci(t *testing.T) {
	test := []struct {
		input    int
		expected int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{6, 8},
		{7, 13},
		{8, 21},
		{9, 34},
		{10, 55},
	}

	for _, test := range test {
		if result := Fibonacci(test.input); result != test.expected {
			t.Errorf("Fibonacci(%d) = %d; want %d", test.input, result, test.expected)
		}
	}
}
