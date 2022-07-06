package main

import (
	"fmt"

	"github.com/biraneves/go-exercises/chap04/examples/ex-04-12/class"
)

func main() {

	class := class.NewClass(10)

	for i := 0; i < class.NumberOfStudents; i++ {

		var result int
		fmt.Print("Enter result (1 = pass, 2 = fail): ")
		fmt.Scanln(&result)

		if result == 1 {
			class.Passes++
		} else {
			class.Failures++
		}

	}

	fmt.Printf("Passed: %d\nFailed: %d\n", class.Passes, class.Failures)

	if class.Passes > 8 {

		fmt.Println("Bonus to instructor!")

	}

}
