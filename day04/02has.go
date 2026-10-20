// День 4, задача 2.
// Напиши функцию Has(m map[string]int, key string) bool, которая возвращает
// true, если ключ есть в map. Ловушка: ключ со значением 0.
package main
 
import "fmt"
 
func main() {
	prices := map[string]int{
		"apple":  25,
		"persik": 90,
		"banana": 67,
	}
	prices["free"] = 0
	fmt.Println(Has(prices, "free"))  // true
	fmt.Println(Has(prices, "mango")) // false
}
 
func Has(m map[string]int, key string) bool {
	_, ok := m[key]
	return ok
}
