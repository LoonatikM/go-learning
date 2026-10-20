//ДЕНЬ 5. Задача 03
//Опиши структуру Counter с полем Value. 
//Напиши метод Inc(), который увеличивает Value на 1.

package main 
import "fmt"

type Book struct {
    Title  string
    Author string
    Pages int
}

type Counter struct {
	Value int

}


func main() {
	b1 := Book{Title: "Inferno", Author: "Dan Brawn", Pages: 389}
	b2 := Book{ Title: "Derek Shtorm", Author: "Kasl", Pages: 215}
	
	fmt.Println(b1.IsLong(), b2.IsLong()) 
	
	c:= Counter{}
	c.Inc()
	c.Inc()
	fmt.Println(c.Value)   // должно быть 2


}

func (b Book)  IsLong() bool {
	 return b.Pages > 300
}

func (c *Counter) Inc() {
	 c.Value++
}