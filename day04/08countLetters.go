// День 4, задача 8.
// CountLetters считает, сколько раз встречается каждая буква в строке.
package main
 
import "fmt"
 
func main() {
	fmt.Println(CountLetters("banana")) // map[97:3 98:1 110:2]
 
	// Ключи печатаются числами, потому что rune — это код символа.
	// Чтобы увидеть буквы:
	for ch, n := range CountLetters("banana") {
		fmt.Println(string(ch), n)
	}
}
 
func CountLetters(s string) map[rune]int {
	counts := make(map[rune]int)
	for _, ch := range s {
		counts[ch]++
	}
	return counts
}
 