package main

import (
	"errors"
	"fmt"
)

var ErrNoPhone = errors.New("phone is empty")

type Speaker interface {
	Speak() string
}

type Dog struct {
	Name string
}

type Cat struct {
	Name string
}

func (d Dog) Speak() string {
	return "Гав! Я " + d.Name
}

func (c Cat) Speak() string {
	return "Мяу! Я " + c.Name
}

type Notifier interface {
	Send(msg string) error
}

type EmailNotifier struct {
	Address string
}

func (e EmailNotifier) Send(msg string) error {
	fmt.Printf("email на %s: %s\n", e.Address, msg)
	return nil

}

type SMSNotifier struct {
	Phone string
}

func (s SMSNotifier) Send(msg string) error {
	if s.Phone == "" {
		return ErrNoPhone
	}

	fmt.Printf("sms на %s: %s\n", s.Phone, msg)
	return nil
}

func NotifyAll(ns []Notifier, msg string) {
	for _, n := range ns {
		err := n.Send(msg)
		if err != nil {
			fmt.Printf("не удалось отправить: %v\n", err)
		}

	}
}

func main() {
	speakers := []Speaker{Dog{"Рекс"}, Cat{"Мурка"}}
	for _, s := range speakers {
		fmt.Println(s.Speak())

	}

	notifiers := []Notifier{
		EmailNotifier{Address: "anya@mail.com"},
		SMSNotifier{Phone: "87773456691"},
		SMSNotifier{Phone: ""}}

	NotifyAll(notifiers, "Заказ готов")

}
