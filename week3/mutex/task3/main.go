package main

import (
	"fmt"
	"sync"
)

type PageStats struct {
	mu     sync.RWMutex
	visits map[string]int
}

func NewPageStats() *PageStats {
	return &PageStats{visits: make(map[string]int)}

}

func (p *PageStats) Visit(page string) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.visits[page]++

}

func (p *PageStats) Count(page string) int {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return p.visits[page]

}

func main() {
	stats := NewPageStats()
	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			stats.Visit("/home")

		}()

	}
	for i := 0; i < 30; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			stats.Visit("/about")

		}()
	}

	wg.Wait()
	fmt.Println("/home:", stats.Count("/home"))
	fmt.Println("/about:", stats.Count("/about"))

}
