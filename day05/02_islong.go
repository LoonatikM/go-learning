// ДЕНЬ 5. Задача 02
//Добавь к Book метод IsLong() bool, который возвращает true, 
//если страниц больше 300. Проверь на обеих книгах.

package main 
import "fmt"

type Book struct {
    Title  string
    Author string
    Pages int
}


func main() {
	b1 := Book{Title: "Inferno", Author: "Dan Brawn", Pages: 389}
	b2 := Book{ Title: "Derek Shtorm", Author: "Kasl", Pages: 215}
	
	fmt.Println(b1.IsLong(), b2.IsLong())   

}

func (b Book)  IsLong() bool {
	 return b.Pages > 300
    
}

