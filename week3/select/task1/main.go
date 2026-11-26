package main

import (
	"fmt"
	"time"
)

func main() {
	chA := make(chan string)
	chB := make(chan string)

	go func() {
		time.Sleep(500 * time.Millisecond)
		chA <- "A"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		chB <- "B"
	}()

	select {
	case msg := <-chA:
		fmt.Println("первым пришел", msg)
	case msg := <-chB:
		fmt.Println("первым пришел", msg)

	}

}
