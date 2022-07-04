package healthProfile

import (
	"fmt"

	sv "github.com/biraneves/screen-visual"
)

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

func (hp *HealthProfile) SetBirthMonth(birthMonth int) {

	if birthMonth >= 1 && birthMonth <= 12 {

		hp.birthMonth = birthMonth

	}

}

func (hp *HealthProfile) GetBirthMonth() int {

	return hp.birthMonth

}

func (hp *HealthProfile) SetBirthYear(birthYear int) {

	if birthYear >= 1900 && birthYear <= 2004 {

		hp.birthYear = birthYear

	}

}

func (hp *HealthProfile) GetBirthDate() string {

	s := fmt.Sprintf("%d/%d/%d", hp.birthDay, hp.birthMonth, hp.birthYear)
	return s

}

func (hp *HealthProfile) SetHeight(height float64) {

	if height > 0.0 {

		hp.height = height

	}

}

func (hp *HealthProfile) GetHeight() float64 {

	return hp.height

}

func (hp *HealthProfile) SetWeight(weight float64) {

	if weight > 0.0 {

		hp.weight = weight

	}

}

func (hp *HealthProfile) GetWeight() float64 {

	return hp.weight

}

func (hp *HealthProfile) GetAge(currentYear int) int {

	return currentYear - hp.birthYear

}

func (hp *HealthProfile) GetMaxHeartRate(currentYear int) int {

	return 220 - hp.GetAge(currentYear)

}

func (hp *HealthProfile) GetTargetHeartRate(currentYear int) (int, int) {

	min := int(0.50 * float64(hp.GetMaxHeartRate(currentYear)))
	max := int(0.85 * float64(hp.GetMaxHeartRate(currentYear)))

	return min, max

}

func (hp *HealthProfile) GetBMI() float64 {

	return hp.weight / (hp.height * hp.height)

}

func GetBMITable() string {

	s := fmt.Sprintf("%s\n", sv.SLine("=", 42))
	s += "BMI (kg/m^2)     Status\n"
	s += fmt.Sprintf("%s\n", sv.SLine("-", 42))
	s += "< 16.0           Severe thinness\n"
	s += "16.0 - 17.0      Moderate thinness\n"
	s += "17 - 18.5        Mild thinness\n"
	s += "18.5 - 25.0      Normal\n"
	s += "25.0 - 30.0      Overweight\n"
	s += "30.0 - 35.0      Obese class I\n"
	s += "35.0 - 40.0      Obese class II\n"
	s += ">40.0            Obese class III\n"
	s += fmt.Sprintf("%s\n", sv.SLine("=", 42))

	return s

}