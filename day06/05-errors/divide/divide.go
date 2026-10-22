//День 6. Задача 1
//делит a на b;
//если b == 0, возвращает ошибку "division by zero".

package main

import (
	"errors"
	"fmt"
)

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil

}

func main() {

	r, err := Divide(10, 2)
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println(r)
	}

	r, err = Divide(10, 0)
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println(r)
	}

}
