//День 6. Повторение
//Напиши функцию, которая получает slice чисел и
//возвращает новый slice без повторов. Порядок
//сохраняется: каждое число остаётся там, где
//встретилось в первый раз.

package main

import "fmt"

func Unique(nums []int) []int {
	seen := map[int]bool{}
	result := []int{}

	for _, n := range nums {
		// 1. если n уже есть в seen — пропусти его
		if seen[n] {
			continue
		}
		seen[n] = true
		result = append(result, n)

	}
	return result
}

func main() {
	fmt.Println(Unique([]int{3, 1, 3, 2, 1})) // [3 1 2]
	fmt.Println(Unique([]int{5, 5, 5}))       // [5]
	fmt.Println(Unique([]int{1, 2, 3}))       // [1 2 3]
	fmt.Println(Unique([]int{}))              // []
}
