package class

type Class struct {
	grades []int
}

func (c *Class) NewGrade(grade int) {

	if grade >= 0 && grade <= 100 {

		c.grades = append(c.grades, grade)

	}

}

func (c *Class) GetGradeCount() int {

	return len(c.grades)

}

func (c *Class) GetTotal() int {

	total := 0

	if len(c.grades) > 0 {

		for _, v := range c.grades {

			total += v

		}

	}

	return total

}

func (c *Class) GetAverage() float64 {

	average := -1.0

	if len(c.grades) > 0 {

		average = float64(c.GetTotal()) / float64(len(c.grades))

	}

	return average

}