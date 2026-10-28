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
	lib.Add("SQL", 1)

	tryTake(lib, "Go")
	tryTake(lib, "SQL")
	tryTake(lib, "SQL")
	tryTake(lib, "Python")
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
}
