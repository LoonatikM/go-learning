package library

import (
	"errors"
)

var ErrBookNotFound = errors.New("book not found")
var ErrNoCopies = errors.New("no copies left")

type Library struct {
	books map[string]int // название → сколько экземпляров
}

func NewLibrary() *Library {
	return &Library{books: make(map[string]int)}

}
func (l *Library) Add(title string, copies int) {
	l.books[title] = l.books[title] + copies

}
func (l *Library) Take(title string) error {
	count, ok := l.books[title]
	if !ok {
		return ErrBookNotFound
	}
	if count == 0 {
		return ErrNoCopies
	}
	l.books[title] -= 1 // уменьши на 1
	return nil
}
func (l *Library) Count(title string) int {
	return l.books[title]

}
