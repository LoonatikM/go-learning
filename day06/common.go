package main

import "fmt"

func Common(a, b []int) []int {
	inB := map[int]bool{}
	added := map[int]bool{}
	result := []int{}

	// Цикл 1: запоминаем все числа из b
	for _, x := range b {
		inB[x] = true
	}

	// Цикл 2: идём по a
	for _, x := range a {
		if inB[x] && !added[x] {
			added[x] = true
			result = append(result, x)
		}

	}

	return result
}

func main() {
	fmt.Println(Common([]int{1, 2, 3, 2, 4}, []int{2, 4, 4, 5}))
	fmt.Println(Common([]int{1, 2}, []int{3, 4}))
}
