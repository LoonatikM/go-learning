package main

import (
	"fmt"
	"sync"
)

var once sync.Once

func connectDB() {
	once.Do(func() {
		fmt.Println("подключаюсь к базе данных")
	})

}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			connectDB()
			fmt.Printf("goroutine %d работает\n", n)

		}(i)

	}
	wg.Wait()
}
