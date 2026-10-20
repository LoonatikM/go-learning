//// День 6. Повторение
//Верни первое число, которое встретилось во второй раз.
//Если повторов нет, верни -1.

package main

import "fmt"

func FirstRepeat(nums []int) int {
	seen := map[int]bool{}

	for _, n := range nums {
		if seen[n] {
			return n
		}
		seen[n] = true
	}
	return -1
}

func main() {
	fmt.Println(FirstRepeat([]int{2, 5, 1, 5, 2}))
	fmt.Println(FirstRepeat([]int{1, 2, 3}))
	fmt.Println(FirstRepeat([]int{}))
}
