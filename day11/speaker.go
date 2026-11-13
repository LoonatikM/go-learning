package main

import "fmt"

type Speaker interface {
	Speak() string
}

type Cat struct {
	Name string
}

type Dog struct {
	Name  string
	Breed string
}

func (d Dog) Speak() string {
	return d.Name + ": гав"
}

func (c Cat) Speak() string {
	return c.Name + ": мяу"
}

func Info(s Speaker) {
	fmt.Println(s.Speak())
	n, ok := s.(Dog)
	if ok {
		fmt.Println("порода:", n.Breed)
	}

}

func main() {
	animals := []Speaker{
		Cat{Name: "барсик"},
		Dog{Name: "Зубик", Breed: "самоед"},
	}

	for _, a := range animals {
		Info(a)
	}

}
