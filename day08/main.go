package main

import (
	"fmt"
	"go-learning/day08/groupbylength"
)

func main() {
	// map[2:[go is] 3:[map fun] 5:[slice]]
	fmt.Println(groupbylength.GroupByLength([]string{"go", "map", "is", "fun", "slice"}))
}
