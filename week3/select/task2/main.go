package main

import (
	"errors"
	"fmt"
	"time"
)

var ErrTimeout = errors.New("timeout")

func slowService(delay time.Duration) chan string {
	ch := make(chan string, 1)

	go func() {
		time.Sleep(delay)
		ch <- "Данные"

	}()
	return ch

}

func CallWithTimeout(delay, timeout time.Duration) (string, error) {
	ch := slowService(delay)

	select {
	case msg := <-ch:
		return msg, nil
	case <-time.After(timeout):
		return "", ErrTimeout

	}
}

func main() {

	d, err := CallWithTimeout(1*time.Second, 2*time.Second)
	if errors.Is(err, ErrTimeout) {
		fmt.Println("сервис не ответил вовремя") // напечатать "сервис не ответил вовремя"
	} else {
		fmt.Println("ответ:", d) // напечатать "ответ:" и d
	}

	d, err = CallWithTimeout(3*time.Second, 2*time.Second) // не успеет
	if errors.Is(err, ErrTimeout) {
		fmt.Println("сервис не ответил вовремя") // напечатать "сервис не ответил вовремя"
	} else {
		fmt.Println("ответ:", d) // напечатать "ответ:" и d
	}

}
