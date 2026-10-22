//День 6. Задача 2
//Превращает строку в возраст:
//строка не число → вернуть ошибку;
//возраст меньше 0 или больше 150 → ошибка с самим числом: "invalid age: 200";
//иначе вернуть возраст и nil

package main

import (
	"fmt"
	"strconv"
)

func ParseAge(s string) (int, error) {
	n, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	if n < 0 || n > 150 {
		return 0, fmt.Errorf("invalid age: %d", n)
	}

	return n, nil
}

func main() {
	inputs := []string{"25", "abc", "-3", "200"}

	for _, s := range inputs {
		age, err := ParseAge(s)
		if err != nil {
			fmt.Println(s, "error:", err)
		} else {
			fmt.Println(s, "age:", age)

		}
	}
}
