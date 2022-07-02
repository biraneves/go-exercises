package main

import (
	"fmt"
	"github.com/biraneves/go-exercises/chap03/exercises/ex-03-11/account"
	sv "github.com/biraneves/screen-visual"
)

func main() {

	account1 := account.NewAccount("Jane Green", 200.0)
	account2 := account.NewAccount("John Blue", -18.34)

	fmt.Printf("%s's initial balance: $%.2f\n", account1.GetName(), account1.GetBalance())
	fmt.Printf("%s's initial balance: $%.2f\n", account2.GetName(), account2.GetBalance())

	sv.Lineln("-", 64)

	var depositAmount float64
	fmt.Print("Inform deposit amount for account1: ")
	fmt.Scanln(&depositAmount)
	account1.Deposit(depositAmount)

	sv.Lineln("-", 64)

	fmt.Printf("%s's balance: $%.2f\n", account1.GetName(), account1.GetBalance())
	fmt.Printf("%s's balance: $%.2f\n", account2.GetName(), account2.GetBalance())

	sv.Lineln("-", 64)

	fmt.Println("Inform deposit amount for account2: ")
	fmt.Scanln(&depositAmount)
	account2.Deposit(depositAmount)

	sv.Lineln("-", 64)

	var amount float64
	fmt.Print("Inform withdrawal amount for account1: ")
	fmt.Scanln(&amount)
	account1.Withdraw(amount)
	fmt.Print("Inform withdrawal amount for account2: ")
	fmt.Scanln(&amount)
	account2.Withdraw(amount)

	sv.Lineln("-", 64)

	fmt.Printf("%s's balance: $%.2f\n", account1.GetName(), account1.GetBalance())
	fmt.Printf("%s's balance: $%.2f\n", account2.GetName(), account2.GetBalance())

}