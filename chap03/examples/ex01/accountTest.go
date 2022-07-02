package main

import (
	"fmt"
	sv "github.com/biraneves/screen-visual"
	"github.com/biraneves/go-exercises/chap03/examples/ex01/account"
)

func main() {

	var account account.Account
	var name string

	sv.Lineln("=", 64)
	fmt.Println("Account initial name:", account.GetName())
	sv.Lineln("-", 64)
	fmt.Print("Your name: ")
	fmt.Scanln(&name)
	account.SetName(name)
	fmt.Println("Account name:", account.GetName())
	sv.Lineln("=", 64)

}