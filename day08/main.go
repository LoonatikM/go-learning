package main

import (
	"errors"
	"fmt"
	"go-learning/day08/groupbylength"
	"go-learning/day08/library"
)

func main() {
	// map[2:[go is] 3:[map fun] 5:[slice]]
	fmt.Println(groupbylength.GroupByLength([]string{"go", "map", "is", "fun", "slice"}))

	lib := library.NewLibrary()
	lib.Add("Go", 2)
	lib.Add("Go", 1)
	fmt.Println("Go:", lib.Count("Go")) // 3

	for i := 1; i <= 4; i++ {
		err := lib.Take("Go")
		if errors.Is(err, library.ErrNoCopies) {
			fmt.Println("попытка", i, ": книга закончилась")
		} else if err != nil {
			fmt.Println("попытка", i, ": другая ошибка:", err)
		} else {
			fmt.Println("попытка", i, ": взяла, осталось", lib.Count("Go"))
		}
	}

	err := lib.Take("Python")
	if errors.Is(err, library.ErrBookNotFound) {
		fmt.Println("Python: такой книги нет")
	}
	fmt.Println("Python:", lib.Count("Python")) // 0
}
