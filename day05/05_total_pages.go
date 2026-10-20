// ДЕНЬ 5. Задача 5

package main 
import "fmt"

type Book struct {
	Title string
	Author string
	Pages int
}

func main() {
	books := []Book{
    {Title: "Inferno", Author: "Dan Brown", Pages: 389},
    {Title: "Derek Storm", Author: "Castle", Pages: 215},
    {Title: "Go", Author: "Donovan", Pages: 380},
}
fmt.Println(TotalPages(books))   // 984

}

func TotalPages(books []Book) int{
	sum:=0
	for _, page:= range books {
		sum = sum+page.Pages
	}
	return sum
}