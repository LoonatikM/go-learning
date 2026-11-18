package main

import "fmt"

func main() {
	ch := make(chan string, 3)
	ch <- "раз"
	ch <- "два"
	ch <- "три"
	fmt.Println(<-ch, <-ch, <-ch)

}
