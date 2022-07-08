package main

import (
	"fmt"
	"github.com/biraneves/go-exercises/chap03/examples/ex02/account"
)

func main() {

	account1 := account.New("Jane Green")
	account2 := account.New("John Blue")

	fmt.Println("Account 1 name:", account1.Name)
	fmt.Println("Account 2 name:", account2.Name)

}