package main

import "fmt"

func ContainsDuplicate(nums []int) bool {
	seen := map[int]bool{}
	for _, n := range nums {
		if seen[n] {
			return true

		}
		seen[n] = true

	}
	return false
}

func main() {
	fmt.Println(ContainsDuplicate([]int{1, 2, 3, 1})) // true  (1 дважды)
	fmt.Println(ContainsDuplicate([]int{1, 2, 3, 4})) // false
	fmt.Println(ContainsDuplicate([]int{}))           // false

}
