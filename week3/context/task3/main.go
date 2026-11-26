package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Result struct {
	URL string
	Err error
}

func fetch(ctx context.Context, url string, delay time.Duration) Result {
	select {
	case <-time.After(delay):
		return Result{URL: url, Err: nil}
	case <-ctx.Done():
		return Result{URL: url, Err: ctx.Err()}

	}

}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2500*time.Millisecond)
	defer cancel()
	sites := []struct {
		url   string
		delay time.Duration
	}{
		{"site-a.com", 1 * time.Second},
		{"site-b.com", 2 * time.Second},
		{"site-c.com", 3 * time.Second},
	}
	results := make(chan Result)
	var wg sync.WaitGroup

	for _, s := range sites {
		wg.Add(1)
		go func(u string, d time.Duration) {
			defer wg.Done()
			results <- fetch(ctx, u, d)

		}(s.url, s.delay)
	}
	go func() {
		wg.Wait()      // дождаться всех
		close(results) // закрыть канал
	}()

	for r := range results {
		if r.Err != nil {
			fmt.Printf("отменено: %s (%v)\n", r.URL, r.Err)
		} else {
			fmt.Println("загружено:", r.URL)
		}
	}

}
