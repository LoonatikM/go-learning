package main

import (
	"errors"
	"fmt"
)

var ErrEmpty = errors.New("empty name")

func Greet(name string) (string, error) {
	if name == "" {
		return "", ErrEmpty
	}
	return "Hello, " + name, nil

}

func main() {
	defer fmt.Println("конец")

	msg, err := Greet("")
	if err != nil {
		fmt.Println("ошибка:", err)
	} else {
		fmt.Println(msg)
	}

	msg, err = Greet("Go")
	if err != nil {
		fmt.Println("ошибка:", err)
	} else {
		fmt.Println(msg)
	}

}
