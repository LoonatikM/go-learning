package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) Introduce() string {
	return fmt.Sprintf("Меня зовут %s, мне %d лет", p.Name, p.Age)
}

type Employee struct {
	Person
	Company string
}

func (e Employee) Introduce() string {
	return fmt.Sprintf("%s. Работаю в %s", e.Person.Introduce(), e.Company)
}

type Introducer interface {
	Introduce() string
}

func main() {

	e := []Introducer{
		Person{Name: "Олег", Age: 30},
		Employee{
			Person:  Person{Name: "Аня", Age: 25},
			Company: "Google",
		},
	}

	for _, r := range e {
		fmt.Println(r.Introduce())

	}

}
