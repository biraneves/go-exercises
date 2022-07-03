package date

import "fmt"

type Date struct {
	day, month, year int
}

func NewDate(day, month, year int) *Date {

	d := new(Date)

	d.day = day
	d.month = month
	d.year = year

	return d

}

func (d *Date) SetDay(day int) {

	d.day = day

}

func (d *Date) GetDay() int {

	return d.day

}

func (d *Date) SetMonth(month int) {

	d.month = month

}

func (d *Date) GetMonth() int {

	return d.month

}

func (d *Date) SetYear(year int) {

	d.year = year

}

func (d *Date) GetYear() int {

	return d.year

}

func (d *Date) DisplayDate() string {

	s := fmt.Sprintf("%d/%d/%d", d.day, d.month, d.year)

	return s

}