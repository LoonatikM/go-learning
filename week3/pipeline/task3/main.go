package main

import (
	"fmt"
	"sync"
	"time"
)

func generate(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func slowSquare(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			time.Sleep(500 * time.Millisecond)
			out <- n * n
		}
		close(out)
	}()
	return out
}

func merge(channels ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup

	for _, c := range channels {
		wg.Add(1)
		go func(ch <-chan int) {
			defer wg.Done()
			for v := range ch {
				out <- v

			}
		}(c)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func main() {
	start := time.Now()

	for s := range slowSquare(generate(1, 2, 3, 4, 5, 6)) {
		fmt.Println(s)
	}
	fmt.Println("один этап:", time.Since(start))

	start = time.Now()

	nums := generate(1, 2, 3, 4, 5, 6)
	c1 := slowSquare(nums)
	c2 := slowSquare(nums)
	c3 := slowSquare(nums)

	for s := range merge(c1, c2, c3) {
		fmt.Println(s)
	}

	fmt.Println("три этап:", time.Since(start))

}
