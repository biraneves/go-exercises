package main

import (
	"fmt"
	"github.com/biraneves/go-exercises/chap03/exercises/ex-03-14/date"
)

func main() {

	date1 := date.NewDate(3, 7, 2022)
	fmt.Println("Date:", date1.DisplayDate())

	date1.SetDay(23)
	date1.SetMonth(4)
	date1.SetYear(1973)
	fmt.Println("Birth date:", date1.DisplayDate())

}