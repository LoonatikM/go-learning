package main

import "fmt"

func main() {
	ch := make(chan int)

	go func() {
		for i := 1; i <= 5; i++ {
			ch <- i * i
		}
		close(ch)
	}()
	sum := 0

	for n := range ch {
		fmt.Println(n)
		sum += n
	}
	fmt.Println("sum:", sum)
	fmt.Println("chanel close")

}
