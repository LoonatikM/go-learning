package main

import (
	"fmt"
	"sync"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("worker %d начал\n", id)
	fmt.Printf("worker %d закончил", id)

}

func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, &wg)

	}
	wg.Wait()
	fmt.Println("готово")

}
