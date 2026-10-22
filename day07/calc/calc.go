package calc

import "errors"

func Add(a, b int) int {
	return a + b
}

func Sub(a, b int) int {
	return a - b
}
func Multiply(a, b int) int {
	return a * b
}

func Divide(a, b int) (int, error) {
	if isZero(b) {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

func isZero(x int) bool {
	return x == 0
}
