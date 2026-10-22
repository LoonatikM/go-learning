// День 6. Задача 5

package main

import (
	"errors"
	"fmt"
)

var ErrUserNotFound = errors.New("user not found")

func FindUser(users map[int]string, id int) (string, error) {
	name, ok := users[id]
	if ok {
		return name, nil
	}
	return "", ErrUserNotFound
}

func main() {
	users := map[int]string{1: "Aida", 2: "Dana"}

	ids := []int{1, 10}
	for _, id := range ids {
		name, err := FindUser(users, id)
		if err != nil {
			fmt.Println(id, "error:", err)
		} else {
			fmt.Println(id, "name:", name)

		}
	}
}
