package heartRate

import "fmt"

type HeartRate struct {
	FirstName, LastName             string
	BirthDay, BirthMonth, BirthYear int
}

func NewHeartRate(firstName, lastName string, birthDay, birthMonth, birthYear int) *HeartRate {

	hr := new(HeartRate)

	hr.FirstName = firstName
	hr.LastName = lastName
	hr.BirthDay = birthDay
	hr.BirthMonth = birthMonth
	hr.BirthYear = birthYear

	return hr

}

func (hr *HeartRate) GetFullName() string {

	s := fmt.Sprintf("%s %s", hr.FirstName, hr.LastName)
	return s

}

func (hr *HeartRate) GetAge(currentYear int) int {

	return currentYear - hr.BirthYear

}

func (hr *HeartRate) GetBirthDate() string {

	s := fmt.Sprintf("%d/%d/%d", hr.BirthDay, hr.BirthMonth, hr.BirthYear)
	return s

}

func (hr *HeartRate) GetMaxHeartRate(currentYear int) int {

	return 220 - hr.GetAge(currentYear)

}

func (hr *HeartRate) GetTargetHeartRates(currentYear int) (int, int) {

	minTargetHeartRate := int(0.50 * float64(hr.GetMaxHeartRate(currentYear)))
	maxTargetHeartRate := int(0.85 * float64(hr.GetMaxHeartRate(currentYear)))

	return minTargetHeartRate, maxTargetHeartRate

}
