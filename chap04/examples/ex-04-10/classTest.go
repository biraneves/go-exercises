package main

import (
	"fmt"

	"github.com/biraneves/go-exercises/chap04/examples/ex-04-10/class"
)

func main() {

	var grade int
	var class class.Class
	
	fmt.Print("Enter grade or -1 to quit: ")
	fmt.Scanln(&grade)

	for grade != -1 {

		class.NewGrade(grade)

		fmt.Print("Enter grade or -1 to quit: ")
		fmt.Scanln(&grade)

	}

	if class.GetGradeCount() != 0 {

		fmt.Printf("\nTotal of the %d grades entered is %d\n", class.GetGradeCount(), class.GetTotal())
		fmt.Printf("Class average is %.2f\n", class.GetAverage())

	} else {

		fmt.Println("No grades entered.")

	}

}