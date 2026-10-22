package main

import (
	"errors"
	"fmt"

	"go-learning/day07/bank"
	"go-learning/day07/calc"
)

func main() {
	fmt.Println(calc.Add(2, 3))      // 5
	fmt.Println(calc.Sub(7, 4))      // 5
	fmt.Println(calc.Multiply(3, 4)) // 5

	r, err := calc.Divide(10, 2)
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println(r)
	}

	r, err = calc.Divide(10, 0)
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println(r)
	}

	acc := bank.NewAccount("Ajgerim")
	// положи 100, сними 500 (обработай через errors.Is), сними 30
	// выведи acc.Balance() — должно быть 70
	acc.Deposit(100)

	err = acc.Withdraw(500)
	if errors.Is(err, bank.ErrInsufficientFunds) {
		fmt.Println("недостаточно средств, баланс:", acc.Balance())
	}

	err = acc.Withdraw(30)
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println("сняли 30, баланс:", acc.Balance())
	}

}
