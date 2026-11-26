package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	mu    sync.Mutex
	value int
}

func (c *Counter) Print() {
	c.mu.Lock()
	defer c.mu.Unlock()
	fmt.Println("значение:", c.value)
}

func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
	fmt.Println("значение:", c.value)
}

func main() {
	c := &Counter{}
	c.Inc()
}
