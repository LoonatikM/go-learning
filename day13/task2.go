package main

import (
	"errors"
	"fmt"
)

var ErrUserNotFound = errors.New("пользователь не найден")

var users = map[int]string{
	1: "Аня",
	2: "Олег",
}

func FindUser(id int) (string, error) {
	name, ok := users[id]
	if !ok {
		return "", ErrUserNotFound
	}
	return name, nil
}

func GetUserName(id int) (string, error) {
	name, err := FindUser(id)
	if err != nil {
		return "", fmt.Errorf("GetUserName id=%d: %w", id, err)
	}
	return name, nil
}

func main() {
	name, err := GetUserName(1) // 1. вызвали, получили имя и ошибку

	if errors.Is(err, ErrUserNotFound) { // 2. это ошибка «не найден»?
		fmt.Println("не найден:", err) //    да → печатаем ошибку
	} else {
		fmt.Println(name) //    нет → печатаем имя
	}
	name, err = GetUserName(5)
	if errors.Is(err, ErrUserNotFound) {
		fmt.Println("не найден:", err)
	} else {
		fmt.Println(name)
	}

}
