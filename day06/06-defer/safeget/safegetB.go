// День 7. Задача 2B
// без паники и без recover: сначала проверь, что индекс подходит;
// индекс подходит, если i >= 0 и i < len(nums);
// если не подходит, верни 0 и fmt.Errorf("index %d out of range", i);
// иначе верни nums[i] и nil.
// В main вызови для индексов 1, 10 и -1, обработай ошибку через if err != nil.

package main

import (
	"fmt"
)

func Get(nums []int, i int) (int, error) {
	if i < 0 || i >= len(nums) {
		return 0, fmt.Errorf("index %d out of range", i)
	}
	return nums[i], nil
}

func main() {
	nums := []int{10, 20, 30}
	indexes := []int{1, 10, -1, 0, 3}

	for _, i := range indexes {
		v, err := Get(nums, i)
		// если err не nil — напечатай "error:" и err
		if err != nil {
			fmt.Println("error:", err)
		} else {
			fmt.Println(v)
		}

	}
}
