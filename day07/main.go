package main

import (
	"fmt"

	"go-learning/day07/calc"
)

func main() {
	fmt.Println(calc.Add(2, 3))      // 5
	fmt.Println(calc.Sub(7, 4))      // 5
	fmt.Println(calc.Multiply(3, 4)) // 5

	r, err := calc.Divide(10, 2)
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println(r)
	}

	r, err = calc.Divide(10, 0)
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println(r)
	}

}
