// День 7. Задача 1
// в самом начале печатает start: <name>;
// с помощью defer печатает end: <name> при любом выходе из функции;
// если name пустая строка, возвращает ошибку "empty name";
// иначе печатает processing <name> и возвращает nil.

package main

import (
	"errors"
	"fmt"
)

func Process(name string) error {

	fmt.Println("start:", name)
	defer fmt.Println("end:", name)

	if name == "" {
		return errors.New("empty name")
	}

	fmt.Println("processing", name)

	return nil

}

func main() {
	inputs := []string{"report", ""}

	for _, s := range inputs {

		err := Process(s)
		if err != nil {
			fmt.Println("error:", err)
		}
	}

}
