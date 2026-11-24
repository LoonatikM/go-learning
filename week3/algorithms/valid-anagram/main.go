package main

import "fmt"

func IsAnagram(s, t string) bool {
	count := map[rune]int{}

	if len(s) != len(t) {
		return false
	}
	for _, ch := range s {

		count[ch]++
	}
	for _, ch := range t {
		count[ch]--

	}
	for _, v := range count {
		if v != 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(IsAnagram("listen", "silent")) // true
	fmt.Println(IsAnagram("rat", "car"))       // false
	fmt.Println(IsAnagram("aab", "abb"))       // false
}
