// День 4, задача 5.
// CountExpensive возвращает количество товаров дороже limit (строго больше).
package main
 
import "fmt"
 
func main() {
	prices := map[string]int{
		"apple":  25,
		"persik": 90,
		"banana": 67,
	}
	prices["free"] = 0
	fmt.Println(CountExpensive(prices, 50)) // 2
	fmt.Println(CountExpensive(prices, 67)) // 1, проверка на границе
}
 
func CountExpensive(m map[string]int, limit int) int {
	sum := 0
	for _, price := range m {
		if limit < price {
			sum++
		}
	}
	return sum
}
