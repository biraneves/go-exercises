package main

import (
	"fmt"

	"github.com/biraneves/go-exercises/chap04/examples/ex-04-08/class"
)

func main() {

	var class1 class.Class

	for i := 0; i < 10; i++ {

		fmt.Print("Enter grade: ")
		fmt.Scanln(&class1.Grades[i])

	}

	fmt.Printf("\nTotal of all 10 grades is %d\n", class1.GetTotal())
	fmt.Printf("Class average is %d\n", class1.GetAverage())

}