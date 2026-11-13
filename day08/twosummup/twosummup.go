package twosummup

func TwoSumMap(nums []int, target int) (int, int) {
	seen := map[int]int{} // число → индекс

	for i, n := range nums {
		t := target - n
		value, ok := seen[t]
		if ok {
			return value, i
		}
		seen[n] = i

	}
	return -1, -1
}
