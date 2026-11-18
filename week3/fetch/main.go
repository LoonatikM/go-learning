package main

import (
	"fmt"
	"sync"
	"time"
)

func fetch(url string) {
	time.Sleep(1 * time.Second) // «загрузка»
	fmt.Println("загружено:", url)
}

func main() {
	urls := []string{"site-a.com", "site-b.com", "site-c.com"}

	// вариант 1: по очереди
	start := time.Now()
	for _, u := range urls {
		fetch(u)
	}
	fmt.Println("по очереди:", time.Since(start))

	// вариант 2: параллельно
	var wg sync.WaitGroup
	start = time.Now()
	for _, u := range urls {
		wg.Add(1) // Add
		go func(u string) {
			defer wg.Done()
			fetch(u)
		}(u) // что передать?
	}
	wg.Wait() // дождаться всех
	fmt.Println("параллельно:", time.Since(start))
}
