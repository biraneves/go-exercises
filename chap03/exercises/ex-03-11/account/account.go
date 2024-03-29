package account

import (
	"fmt"
)

type Account struct {
	name    string
	balance float64
}

func NewAccount(name string, balance float64) *Account {

	a := new(Account)

	a.name = name
	if balance > 0.0 {
		a.balance = balance
	}

	return a

}

func (a *Account) Deposit(depositAmount float64) {

	if depositAmount > 0.0 {
		a.balance += depositAmount
	}

}

func (a *Account) Withdraw(amount float64) {

	if amount > a.balance {
		fmt.Println("Withdrawal amount exceeded account balance.")
	} else {
		a.balance -= amount
	}

}

func (a *Account) GetBalance() float64 {

	return a.balance

}

func (a *Account) SetName(name string) {

	a.name = name

}

func (a *Account) GetName() string {

	return a.name

}