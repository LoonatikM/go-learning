// День 4, задача 6.
// MostExpensive возвращает название самого дорогого товара.
package main
 
import "fmt"
 
func main() {
	prices := map[string]int{
		"apple":  25,
		"persik": 90,
		"banana": 67,
	}
	prices["free"] = 0
	fmt.Println(MostExpensive(prices)) // persik
}
 
func MostExpensive(m map[string]int) string {
	sum := 0
	name1 := ""
	for name, price := range m {
		if sum < price {
			sum = price
			name1 = name
		}
	}
	return name1
}
