package main

import (
	"fmt"

	"github.com/biraneves/go-exercises/chap03/exercises/ex-03-15/account"
	sv "github.com/biraneves/screen-visual"
)

func main() {

	account1 := account.NewAccount("Jane Green", 200.00)
	account2 := account.NewAccount("John Blue", -18.34)

	sv.Lineln("-", 64)

	var depositAmount float64
	fmt.Print("Inform deposit amount for account1: ")
	fmt.Scanln(&depositAmount)
	account1.Deposit(depositAmount)

	sv.Lineln("-", 64)

	fmt.Println(account1.DisplayAccount())
	fmt.Println(account2.DisplayAccount())

	sv.Lineln("-", 64)

	fmt.Print("Inform deposit amount for account2: ")
	fmt.Scanln(&depositAmount)
	account2.Deposit(depositAmount)

	sv.Lineln("-", 64)

	fmt.Println(account1.DisplayAccount())
	fmt.Println(account2.DisplayAccount())

	sv.Lineln("-", 64)

	var amount float64
	fmt.Print("Inform withdrawal amount for account1: ")
	fmt.Scanln(&amount)
	account1.Withdraw(amount)
	fmt.Print("Inform withdrawal amount for account2: ")
	fmt.Scanln(&amount)
	account2.Withdraw(amount)

	sv.Lineln("-", 64)

	fmt.Println(account1.DisplayAccount())
	fmt.Println(account2.DisplayAccount())

}
