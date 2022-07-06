package class

type Class struct {
	NumberOfStudents int
	Passes           int
	Failures         int
}

func NewClass(numberOfStudents int) *Class {

	c := new(Class)

	if numberOfStudents > 0 {

		c.NumberOfStudents = numberOfStudents

	}

	c.Passes = 0
	c.Failures = 0

	return c

}
