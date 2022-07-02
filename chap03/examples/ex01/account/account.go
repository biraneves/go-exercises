package account

type Account struct {
	name string
}

func (a *Account) SetName(name string) {

	a.name = name

}

func (a *Account) GetName() string {

	return a.name

}