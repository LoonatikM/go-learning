// День 4, задача 4.
// Total возвращает сумму всех цен в map.
package main
 
import "fmt"
 
func main() {
	prices := map[string]int{
		"apple":  25,
		"persik": 90,
		"banana": 67,
	}
	prices["free"] = 0
	fmt.Println(Total(prices)) // 182
}
 
func Total(m map[string]int) int {
	sum := 0
	for _, price := range m {
		sum = sum + price
	}
	return sum
}
