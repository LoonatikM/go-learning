package main

import "fmt"

func Index[T comparable](s []T, v T) int {
	for index, x := range s {
		if x == v {
			return index
		}
	}
	return -1

}

func main() {
	fmt.Println(Index([]int{10, 20, 30}, 20))           // 1
	fmt.Println(Index([]string{"go", "sql"}, "docker")) // -1
	fmt.Println(Index([]string{"go", "sql"}, "go"))     // 0
}
