package mathutils

import "errors"

var ErrDiviedByZero = errors.New("cannot divide by zero")

func Multiply(a, b int) int {
	return a * b
}

func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, ErrDiviedByZero
	}
	return (a / b), nil
}
