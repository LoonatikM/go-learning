package main

import (
	"errors"
	"fmt"
	"go-learning/day08/groupbylength"
	"go-learning/day08/library"
	"go-learning/day08/twosum"
	"go-learning/day08/without"
)

func main() {
	// map[2:[go is] 3:[map fun] 5:[slice]]
	fmt.Println(groupbylength.GroupByLength([]string{"go", "map", "is", "fun", "slice"}))

	lib := library.NewLibrary()
	lib.Add("Go", 2)
	lib.Add("SQL", 1)

	tryTake(lib, "Go")
	tryTake(lib, "SQL")
	tryTake(lib, "SQL")
	tryTake(lib, "Python")

	fmt.Println(twosum.TwoSum([]int{2, 7, 11, 15}, 9)) // 0, 1   → потому что 2 + 7 = 9
	fmt.Println(twosum.TwoSum([]int{3, 2, 4}, 6))      // 1, 2   → 2 + 4 = 6
	fmt.Println(twosum.TwoSum([]int{3, 3}, 6))

}

func tryTake(l *library.Library, title string) {
	err := l.Take(title)
	if errors.Is(err, library.ErrBookNotFound) {
		fmt.Printf("книги %q нет в библиотеке\n", title)
	} else if errors.Is(err, library.ErrNoCopies) {
		fmt.Printf("%q закончилась, приходи позже\n", title)
	} else if err != nil {
		fmt.Printf("неизвестная ошибка: %v\n", err)
	} else {
		fmt.Printf("взяла %q, осталось %d\n", title, l.Count(title))
	}

	nums := []int{1, 2, 3, 2, 4}
	res := without.Without(nums, 2)

	fmt.Println("nums without.go", res) // [1 3 4]

}
