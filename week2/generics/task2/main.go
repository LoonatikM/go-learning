package main

import (
	"cmp"
	"fmt"
)

type Number interface {
	int | int64 | float64
}

func Sum[T Number](nums []T) T {
	var total T
	for _, n := range nums {
		total += n
	}
	return total

}

func MaxOf[T cmp.Ordered](s []T) (T, bool) {
	if len(s) == 0 {
		var zero T // нулевое значение типа T
		return zero, false
	}
	best := s[0]
	for _, m := range s {
		if best < m {
			best = m
		}
	}
	return best, true

}

func main() {
	fmt.Println(Sum([]int{1, 2, 3}))      // 6
	fmt.Println(Sum([]float64{1.5, 2.5})) // 4

	m, ok := MaxOf([]int{3, 9, 2})
	fmt.Println(m, ok) // 9 true

	w, ok := MaxOf([]string{"go", "apple", "sql"})
	fmt.Println(w, ok) // sql true

	e, ok := MaxOf([]int{})
	fmt.Println(e, ok) // 0 false
}
