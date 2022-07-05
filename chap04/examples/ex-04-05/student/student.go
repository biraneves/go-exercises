package student

type Student struct {
	Name string
	average float64
}

func NewStudent(name string, average float64) *Student {

	s := new(Student)

	s.Name = name

	if average >= 0.0 && average <= 100.0 {
		s.average = average
	}

	return s

}

func (s *Student) SetAverage(average float64) {

	if average >= 0.0 && average <= 100.0 {

		s.average = average

	}

}

func (s *Student) GetAverage() float64 {

	return s.average

}

func (s *Student) GetLetterGrade() string {

	var letterGrade string

	if s.average >= 90.0 {

		letterGrade = "A"

	} else if s.average >= 80 {

		letterGrade = "B"

	} else if s.average >= 70 {

		letterGrade = "C"

	} else if s.average >= 60 {

		letterGrade = "D"

	} else {

		letterGrade = "F"

	}

	return letterGrade

}