//День 6. Повторение
//Напиши функцию, которая получает slice чисел и
//возвращает новый slice без повторов. Порядок
//сохраняется: каждое число остаётся там, где
//встретилось в первый раз.

package main

import (
	"fmt"
	"strings"
)

func UniqueWords(words []string) []string {

	seen := map[string]bool{}
	result := []string{}

	for _, n := range words {
		lower := strings.ToLower(n)
		if seen[lower] {
			continue
		}
		seen[lower] = true
		result = append(result, n)

	}
	return result
}

func main() {
	fmt.Println(UniqueWords([]string{"Go", "is", "fun", "go", "IS", "cool"}))
}
