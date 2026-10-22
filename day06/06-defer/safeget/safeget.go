// День 7ю Задача 2

package main

import "fmt"

func SafeGet(nums []int, i int) int {
	defer func() {
		// если была паника — поймай её через recover и напечатай
		// "поймали панику:" и её значение
		if r := recover(); r != nil {
			fmt.Println("поймали панику:", r)
		}
	}()

	return nums[i]
}

func main() {
	nums := []int{10, 20, 30}
	fmt.Println(SafeGet(nums, 1))  // 20
	fmt.Println(SafeGet(nums, 10)) // поймали панику: ...  затем 0
	fmt.Println("программа работает дальше")
}
