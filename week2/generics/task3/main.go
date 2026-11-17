package main

import "fmt"

type Queue[T any] struct {
	items []T
}

func (q *Queue[T]) Enqueue(v T) { // добавить в конец
	q.items = append(q.items, v)
}

func (q *Queue[T]) Dequeue() (T, bool) { // забрать из НАЧАЛА
	if len(q.items) == 0 {
		var zero T // нулевое значение типа T
		return zero, false
	}
	first := q.items[0]
	q.items = q.items[1:]
	return first, true
}

func (q *Queue[T]) Len() int { // сколько элементов
	return len(q.items)

}

func main() {
	q := Queue[string]{}
	q.Enqueue("Аня")
	q.Enqueue("Боб")
	q.Enqueue("Вика")
	fmt.Println(q.Len()) // 3

	v, ok := q.Dequeue()
	fmt.Println(v, ok) // Аня true
	v, ok = q.Dequeue()
	fmt.Println(v, ok)   // Боб true
	fmt.Println(q.Len()) // 1
}
