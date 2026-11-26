package main

import "fmt"

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

func filterEven(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			if n%2 == 0 {
				out <- n
			}

		}
		close(out)
	}()
	return out
}

func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func main() {
	nums := generate(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	results := square(filterEven(nums))
	sum := 0

	for r := range results {
		fmt.Println(r)
		sum += r

	}
	fmt.Println("sum:", sum)
}
