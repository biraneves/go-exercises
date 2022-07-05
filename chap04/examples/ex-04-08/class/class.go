package class

type Class struct {
	Grades [10]int
}

func (c *Class) GetTotal() int {

	total := 0
	
	for _, value := range c.Grades {

		total += value

	}

	return total

}

func (c *Class) GetAverage() int {

	return c.GetTotal() / 10

}