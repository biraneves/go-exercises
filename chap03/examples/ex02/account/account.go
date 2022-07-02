package account

type Account struct {
	name string
}

func NewAccount(name string) *Account {

	a := new(Account)
	a.name = name
	return a

}

func (a *Account) SetName(name string) {

	a.name = name

}

func (a *Account) GetName() string {

	return a.name

}