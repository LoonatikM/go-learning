// День 4, задача 7.
// AddStock увеличивает количество товара на складе на amount.
// Если товара ещё нет, он должен появиться. Функция ничего не возвращает.
package main
 
import "fmt"
 
func main() {
	stock := map[string]int{"apple": 3}
	AddStock(stock, "apple", 2)
	AddStock(stock, "kiwi", 7)
	fmt.Println(stock) // map[apple:5 kiwi:7]
}
 
func AddStock(stock map[string]int, name string, amount int) {
	v := stock[name] // ключа нет → вернётся 0 (нулевое значение типа), а не ошибка
	stock[name] = amount + v
}
 
// Замечание ментора: то же самое одной строкой: stock[name] += amount.
//
// Почему stock в main изменился, хотя функция ничего не вернула:
// Go всегда передаёт копию, но переменная map хранит ссылку на таблицу.
// Копируется ссылка, и она ведёт на ту же таблицу, что и оригинал.