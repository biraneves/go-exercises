package account

type Account struct {
	Name string
}

func New(name string) *Account {

	a := new(Account)
	a.Name = name
	return a

}