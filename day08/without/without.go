package without

func Without(nums []int, x int) []int {
	result := make([]int, 0, len(nums))

	for _, n := range nums {
		if n != x {
			result = append(result, n)
		} else {
			continue
		}

	}
	return result

}
