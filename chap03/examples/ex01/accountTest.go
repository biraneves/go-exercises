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
	fmt.Println("Account initial name:", account.Name)
	sv.Lineln("-", 64)
	fmt.Print("Your name: ")
	fmt.Scanln(&name)
	account.Name = name
	fmt.Println("Account name:", account.Name)
	sv.Lineln("=", 64)

}