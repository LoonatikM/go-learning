// ErrInsufficientFunds — экспортируемая ошибка
// ErrInvalidAmount — экспортируемая ошибка
// NewAccount(owner string) *Account — конструктор, баланс 0
// (a *Account) Deposit(amount float64) error
// (a *Account) Withdraw(amount float64) error
// (a *Account) Balance() float64 — вернуть баланс, чтобы его можно было прочитать снаружи

package bank

import "errors"

var ErrInsufficientFunds = errors.New("insufficient funds")
var ErrInvalidAmount = errors.New("invalid amount")

type Account struct {
	Owner   string
	balance float64 // строчная буква — снаружи не видно
}

func NewAccount(owner string) *Account {
	return &Account{Owner: owner}

}

func (a *Account) Deposit(amount float64) error {
	if amount <= 0 {
		return ErrInvalidAmount // одно значение — ошибка
	}
	// увеличь баланс
	a.balance += amount
	return nil // ошибки нет
}

func (a *Account) Withdraw(amount float64) error {
	if amount <= 0 {
		return ErrInvalidAmount // одно значение — ошибка
	}

	if amount > a.balance {
		return ErrInsufficientFunds
	}
	a.balance -= amount
	return nil // ошибки нет
}

func (a *Account) Balance() float64 {
	return a.balance
}
