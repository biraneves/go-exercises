package account

type Account struct {
	Name string
	balance float64
}

func New(name string, balance float64) *Account {

	a := new(Account)
	a.Name = name

	if balance > 0.0 {
		a.balance = balance
	}

	return a

}

func (a *Account) Deposit(depositAmount float64) {

	if depositAmount > 0.0 {
		a.balance += depositAmount
	}

}

func (a *Account) Balance() float64 {

	return a.balance

}