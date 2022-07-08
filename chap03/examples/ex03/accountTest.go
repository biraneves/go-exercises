package main

import (
	"fmt"
	"github.com/biraneves/go-exercises/chap03/examples/ex03/account"
	sv "github.com/biraneves/screen-visual"
)

func main() {

	account1 := account.New("Jane Green", 50.00)
	account2 := account.New("John Blue", -7.53)

	fmt.Printf("%s's account balance: $%.2f\n", account1.Name, account1.Balance())
	fmt.Printf("%s's account balance: $%.2f\n", account2.Name, account2.Balance())
	
	sv.Lineln("-", 64)

	var depositAmount float64
	fmt.Print("Enter deposit amount for account 1: ")
	fmt.Scanln(&depositAmount)
	fmt.Printf("Adding %.2f to account1 balance\n", depositAmount)
	account1.Deposit(depositAmount)

	sv.Lineln("-", 64)

	fmt.Printf("%s's account balance: $%.2f\n", account1.Name, account1.Balance())
	fmt.Printf("%s's account balance: $%.2f\n", account2.Name, account2.Balance())

	sv.Lineln("-", 64)

	fmt.Print("Enter deposit amount for account 2: ")
	fmt.Scanln(&depositAmount)
	fmt.Printf("Adding %.2f to account2 balance\n", depositAmount)
	account2.Deposit(depositAmount)

	sv.Lineln("-", 64)

	fmt.Printf("%s's account balance: $%.2f\n", account1.Name, account1.Balance())
	fmt.Printf("%s's account balance: $%.2f\n", account2.Name, account2.Balance())

}