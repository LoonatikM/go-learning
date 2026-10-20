//ДЕНЬ 5. Задача 4
//4. Account из входного чека: поле Balance, 
//методы Deposit(amount int) и Withdraw(amount int) bool. 
//Withdraw уменьшает баланс и возвращает true, 
//а если денег не хватает, ничего не меняет и возвращает false.

package main
import "fmt"

type Account struct {
	Balance int
}

func main() {
	acc := Account{Balance: 100}
	acc.Deposit(50)
	fmt.Println(acc.Withdraw(500))   // false
	fmt.Println(acc.Withdraw(120))   // true
	fmt.Println(acc.Balance)         // 30

}

func (c *Account) Deposit(amount int) {
	 c.Balance =c.Balance+amount

}

func (c *Account) Withdraw(amount int) bool{
	if amount > c.Balance{
		return false
	} 
	c.Balance-=amount
	return true

}