package main

import (
	"fmt"

	"github.com/biraneves/go-exercises/chap04/examples/ex-04-05/student"
)

func main() {

	account1 := student.NewStudent("Jane Green", 93.5)
	account2 := student.NewStudent("John Blue", 72.75)

	fmt.Printf("%s's letter grade is %s\n", account1.Name, account1.GetLetterGrade())
	fmt.Printf("%s's letter grade is %s\n", account2.Name, account2.GetLetterGrade())

}