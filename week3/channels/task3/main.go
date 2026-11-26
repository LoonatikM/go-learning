package main

import (
	"fmt"
	"sync"
	"time"
)

type Result struct {
	URL  string
	Size int
}

func fetch(url string) Result {
	time.Sleep(1 * time.Second)
	return Result{URL: url, Size: len(url) * 100}
}

func main() {
	urls := []string{"site-a.com", "site-b.com", "site-c.com"}
	results := make(chan Result)

	var wg sync.WaitGroup

	for _, ur := range urls {
		wg.Add(1)
		go func(u string) {
			defer wg.Done()
			results <- fetch(u)
		}(ur)

	}

	go func() {
		wg.Wait()
		close(results)
	}()
	total := 0

	for r := range results {
		fmt.Printf("%s : %d\n", r.URL, r.Size)
		total += r.Size
	}

	fmt.Println("Total:", total)
	fmt.Println("chanel close")

}
