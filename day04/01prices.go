// День 4, задача 1.
// Создай map prices с тремя товарами и ценами. Выведи цену одного товара,
// добавь четвёртый, удали один, выведи длину.
package main
 
import "fmt"
 
func main() {
	prices := map[string]int{
		"apple":  25,
		"persik": 90,
		"banana": 67,
	}
	fmt.Println(prices["banana"]) // 67
 
	prices["vino"] = 73
	delete(prices, "persik")
	fmt.Println(len(prices)) // 3
}
