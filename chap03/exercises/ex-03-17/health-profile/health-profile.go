package healthProfile

import "fmt"

type HealthProfile struct {
	FirstName, LastName, Sex        string
	birthDay, birthMonth, birthYear int
	height, weight                  float64
}

func NewHealthProfile(firstName, lastName, sex string, birthDay, birthMonth, birthYear int, height, weight float64) *HealthProfile {

	hp := new(HealthProfile)

	hp.FirstName = firstName
	hp.LastName = lastName
	hp.Sex = sex
	if birthDay >= 1 && birthDay <= 31 {
		hp.birthDay = birthDay
	}
	if birthMonth >= 1 && birthMonth <= 12 {
		hp.birthMonth = birthMonth
	}
	if birthYear >= 1900 && birthYear <= 2004 {
		hp.birthYear = birthYear
	}
	if height > 0.0 {
		hp.height = height
	}
	if weight > 0.0 {
		hp.weight = weight
	}

	return hp

}

func (hp *HealthProfile) GetFullName() string {

	s := fmt.Sprintf("%s %s", hp.FirstName, hp.LastName)
	return s

}

func (hp *HealthProfile) SetBirthDay(birthDay int) {

	if birthDay >= 1 && birthDay <= 31 {

		hp.birthDay = birthDay

	}

}

func (hp *HealthProfile) GetBirthDay() int {

	return hp.birthDay

}
