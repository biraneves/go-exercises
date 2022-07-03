package main

import (
	"fmt"
	"github.com/biraneves/go-exercises/chap03/exercises/ex-03-12/invoice"
	sv "github.com/biraneves/screen-visual"
)

func main() {

	invoice1 := invoice.NewInvoice("001", "Laptop", 2, 1450.00)
	invoice2 := invoice.NewInvoice("002", "Mouse", 15, 28.80)

	sv.Lineln("=", 80)
	fmt.Printf("Number\t\tDescription\t\tQuantity\tPrice\t\tTotal\n")
	sv.Lineln("-", 80)
	fmt.Printf("%s\t\t%s\t\t\t%d\t\t$%.2f\t$%.2f\n", invoice1.GetNumber(),
		invoice1.GetDescription(), invoice1.GetQuantity(), invoice1.GetPrice(),
		invoice1.GetInvoiceAmount())
	fmt.Printf("%s\t\t%s\t\t\t%d\t\t$%.2f\t\t$%.2f\n", invoice2.GetNumber(),
		invoice2.GetDescription(), invoice2.GetQuantity(), invoice2.GetPrice(),
		invoice2.GetInvoiceAmount())
	sv.Lineln("-", 80)

}