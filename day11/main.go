package main

import "fmt"

func Double(x any) {
	n, ok := x.(int)
	if !ok {
		fmt.Println("не целое число")
		return
	}
	fmt.Println(n * 2)
}

func Describe(x any) string {
	switch v := x.(type) {
	case int:
		return "число"
	case string:
		return "строка"
	case bool:
		return "да/нет"
	case []int:
		return fmt.Sprintf("срез из %d элементов", len(v))
	default:
		return "неизвестный тип"

	}

}

func main() {

	Double(21)
	Double("21")
	Double(2.5)

	fmt.Println(Describe(5))              // число
	fmt.Println(Describe("Go"))           // строка
	fmt.Println(Describe([]int{1, 2, 3})) // срез из 3 элементов
	fmt.Println(Describe(3.14))
}
