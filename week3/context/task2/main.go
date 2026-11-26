package main

import (
	"context"
	"fmt"
	"sync"
)

func worker(ctx context.Context, jobs chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case j := <-jobs:
			fmt.Println("обрабатываю задачу", j)
		case <-ctx.Done():
			fmt.Println("worker остановлен:", ctx.Err())
			return
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	jobs := make(chan int)
	var wg sync.WaitGroup

	wg.Add(1)

	go worker(ctx, jobs, &wg)

	for i := 1; i <= 3; i++ {
		jobs <- i

	}
	cancel()
	wg.Wait()
	fmt.Println("main завершён")

}
