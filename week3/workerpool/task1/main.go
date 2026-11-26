package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs { // берёт задачи, пока канал не закроют
		fmt.Printf("worker %d взял задачу %d\n", id, j)
		time.Sleep(500 * time.Millisecond) // «работа»
		results <- j * j                   // результат
	}
}

func main() {
	const numWorkers = 9
	jobs := make(chan int, 10)
	results := make(chan int, 10)
	var wg sync.WaitGroup
	start := time.Now()

	// 1. запускаем 3 worker'а
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}

	// 2. кладём 9 задач и закрываем очередь
	for j := 1; j <= 9; j++ {
		jobs <- j
	}
	close(jobs)

	// 3. когда все worker'ы закончат — закрываем results
	go func() {
		wg.Wait()
		close(results)
	}()

	// 4. собираем результаты
	for r := range results {
		fmt.Println("результат:", r)
	}
	fmt.Println("время:", time.Since(start))

}
