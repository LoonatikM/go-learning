package twosum

func TwoSum(nums []int, target int) (int, int) {

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return i, j
			}
		}
	}
	return -1, -1
}
