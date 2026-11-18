package main

import (
	"errors"
	"fmt"
)

var ErrUserNotFound = errors.New("пользователь не найден")

// task 1
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

// task 2
type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {

	return fmt.Sprintf("поле %s: %s", e.Field, e.Message)
}

func Validate(name string, age int) error {
	if name == "" {
		return &ValidationError{Field: "name", Message: "не может быть пустым"}
	}
	if age < 0 || age > 150 {
		return &ValidationError{Field: "age", Message: "недопустимое значение"}
	}

	return nil
}

//task 3

func Register(name string, age int) error {

	err := Validate(name, age)
	if err != nil {
		return fmt.Errorf("регистрация: %w", err)
	}
	return nil

}

func main() {
	// task 1
	name, err := GetUserName(1)

	if errors.Is(err, ErrUserNotFound) {
		fmt.Println("не найден:", err)
	} else {
		fmt.Println(name)
	}
	name, err = GetUserName(5)
	if errors.Is(err, ErrUserNotFound) {
		fmt.Println("не найден:", err)
	} else {
		fmt.Println(name)
	}

	// task 2
	fmt.Println(Validate("", 20))
	fmt.Println(Validate("Аня", 200))
	fmt.Println(Validate("Аня", 25))

	//task 3
	err = Register("", 20)
	fmt.Println("ошибка: %w", err)

	var ve *ValidationError
	if errors.As(err, &ve) {
		fmt.Println("неверное поле:", ve.Field)

	}

}
