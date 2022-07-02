package main

import (
	"fmt"
	"github.com/biraneves/go-exercises/chap03/examples/ex02/account"
)

func main() {

	account1 := account.NewAccount("Jane Green")
	account2 := account.NewAccount("John Blue")

	fmt.Println("Account 1 name:", account1.GetName())
	fmt.Println("Account 2 name:", account2.GetName())

}