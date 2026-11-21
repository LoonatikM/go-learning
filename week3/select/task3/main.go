package main

import (
	"fmt"
	"sync"
)

func worker(jobs chan int, done chan bool, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case j := <-jobs:
			fmt.Println("обрабатываю задачу", j)
		case <-done:
			fmt.Println("worker остановлен")
			return
		}
	}
}

func main() {
	jobs := make(chan int)
	done := make(chan bool)
	var wg sync.WaitGroup

	wg.Add(1)

	go worker(jobs, done, &wg)

	for i := 1; i <= 3; i++ {
		jobs <- i

	}
	done <- true
	wg.Wait()
	fmt.Println("main завершён")

}
