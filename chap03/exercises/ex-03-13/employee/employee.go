package employee

type Employee struct {
	firstName       string
	lastName        string
	paymentPerMonth float64
}

func NewEmployee(firstName, lastName string, paymentPerMonth float64) *Employee {

	e := new(Employee)

	e.firstName = firstName
	e.lastName = lastName
	if paymentPerMonth > 0.0 {
		e.paymentPerMonth = paymentPerMonth
	}

	return e

}

func (e *Employee) SetFirstName(firstName string) {

	e.firstName = firstName

}

func (e *Employee) GetFirstName() string {

	return e.firstName

}

func (e *Employee) SetLastName(lastName string) {

	e.lastName = lastName

}

func (e *Employee) GetLastName() string {

	return e.lastName

}

func (e *Employee) SetPaymentPerMonth(paymentPerMonth float64) {

	if paymentPerMonth > 0.0 {

		e.paymentPerMonth = paymentPerMonth

	}

}

func (e *Employee) GetPaymentPerMonth() float64 {

	return e.paymentPerMonth

}
