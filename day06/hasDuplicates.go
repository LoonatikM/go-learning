// День 6. Повторение
//Есть ли в slice хотя бы один повтор?
//Как только нашла повтор, можно сразу return true,
//дальше проверять незачем.

package main

import "fmt"

func HasDuplicates(nums []int) bool {
	seen := map[int]bool{}

	for _, n := range nums {
		// 1. если n уже есть в seen — пропусти его
		if seen[n] {
			return true
		}
		seen[n] = true
	}
	return false

}

func main() {

	fmt.Println(HasDuplicates([]int{1, 2, 3})) // false
	fmt.Println(HasDuplicates([]int{1, 2, 1})) // true
	fmt.Println(HasDuplicates([]int{}))        // false

}
