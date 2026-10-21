//День 6. Задача 3
//В main: положи 100, попробуй снять 500 и через errors.Is выведи понятное сообщение,
//например "недостаточно средств". Потом сними 30 и выведи баланс (должно быть 70).

package main

import (
	"errors"
	"fmt"
)

var ErrInsufficientFunds = errors.New("insufficient funds")

type BankAccount struct {
	Owner   string
	Balance float64
}

// Deposit — пополнить счёт.
// Если amount <= 0, вернуть ошибку "invalid amount" и баланс не менять.
func (a *BankAccount) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("invalid amount") // одно значение — ошибка
	}
	// увеличь баланс
	a.Balance += amount
	return nil // ошибки нет
}

// Withdraw — снять деньги.
// Если amount <= 0 — ошибка "invalid amount".
// Если денег не хватает — вернуть ErrInsufficientFunds, баланс не менять.
func (a *BankAccount) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("invalid amount") // одно значение — ошибка
	}

	if amount > a.Balance {
		return ErrInsufficientFunds
	}
	a.Balance -= amount
	return nil // ошибки нет
}

func main() {
	acc := BankAccount{Owner: "Ajgerim"}
	acc.Deposit(100)

	err := acc.Withdraw(500)
	if errors.Is(err, ErrInsufficientFunds) {
		fmt.Println("недостаточно средств, баланс:", acc.Balance)
	}

	err = acc.Withdraw(30)
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println("сняли 30, баланс:", acc.Balance)
	}
}
