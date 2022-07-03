package main

import (
	"fmt"
	"github.com/biraneves/go-exercises/chap03/exercises/ex-03-13/employee"
	sv "github.com/biraneves/screen-visual"
)

func main() {

	employee1 := employee.NewEmployee("Jane", "Green", 2000.00)
	employee2 := employee.NewEmployee("John", "Blue", -1800.00)

	sv.Lineln("-", 64)
	
	fmt.Printf("%s's initial payment per month: $%.2f\n", employee1.GetFirstName(), employee1.GetPaymentPerMonth())
	fmt.Printf("%s's initial payment per month: $%.2f\n", employee2.GetFirstName(), employee2.GetPaymentPerMonth())
	employee2.SetPaymentPerMonth(1800.00)
	fmt.Printf("Adjusting %s's payment per month... \n", employee2.GetFirstName())
	fmt.Printf("%s's initial payment per month: $%.2f\n", employee1.GetFirstName(), employee1.GetPaymentPerMonth())
	fmt.Printf("%s's initial payment per month: $%.2f\n", employee2.GetFirstName(), employee2.GetPaymentPerMonth())
	
	sv.Lineln("-", 64)

	fmt.Printf("Raising %s's payment per month in 10%%\n", employee1.GetFirstName())
	employee1.SetPaymentPerMonth(employee1.GetPaymentPerMonth() * 1.10)
	fmt.Printf("Raising %s's payment per month in 10%%\n", employee2.GetFirstName())
	employee2.SetPaymentPerMonth(employee2.GetPaymentPerMonth() * 1.10)
	fmt.Printf("%s's new payment per month: $%.2f\n", employee1.GetFirstName(), employee1.GetPaymentPerMonth())
	fmt.Printf("%s's new payment per month: $%.2f\n", employee2.GetFirstName(), employee2.GetPaymentPerMonth())
	
	sv.Lineln("-", 64)

	fmt.Printf("%s's annual payment: $%.2f\n", employee1.GetFirstName(), employee1.GetPaymentPerMonth() * float64(12))
	fmt.Printf("%s's annual payment: $%.2f\n", employee2.GetFirstName(), employee2.GetPaymentPerMonth() * float64(12))
	
}