//День 5. Задача 01
//Опиши структуру Book с полями Title, Author, Pages. Создай в main две книги, 
//выведи название первой и количество страниц второй, затем измени количество 
//страниц у первой и выведи её целиком.
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
	fmt.Println(b1.Title, b2.Pages)
	b1.Pages = 710
	fmt.Printf("%+v\n", b1)   // {Title:Inferno Author:Dan Brawn Pages:710}
}
