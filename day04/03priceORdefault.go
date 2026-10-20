// День 4, задача 3.
// PriceOrDefault возвращает цену товара, а если его нет, то -1.
package main
 
import "fmt"
 
func main() {
	prices := map[string]int{
		"apple":  25,
		"persik": 90,
		"banana": 67,
	}
	prices["free"] = 0
	fmt.Println(PriceOrDefault(prices, "apple")) // 25
	fmt.Println(PriceOrDefault(prices, "mango")) // -1
}
 
func PriceOrDefault(m map[string]int, key string) int {
	_, ok := m[key]
	if ok == false {
		return -1
	} else {
		return m[key]
	}
}
