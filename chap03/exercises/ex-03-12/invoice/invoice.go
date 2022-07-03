package invoice

type Invoice struct {
	number      string
	description string
	quantity    int
	price       float64
}

func NewInvoice(number, description string, quantity int, price float64) *Invoice {

	i := new(Invoice)
	i.number = number
	i.description = description
	if quantity > 0 {
		i.quantity = quantity
	}
	if price > 0.0 {
		i.price = price
	}

	return i

}

func (i *Invoice) GetNumber() string {

	return i.number

}

func (i *Invoice) SetDescription(description string) {

	i.description = description

}

func (i *Invoice) GetDescription() string {

	return i.description

}

func (i *Invoice) SetQuantity(quantity int) {

	if quantity > 0 {
		i.quantity = quantity
	}

}

func (i *Invoice) GetQuantity() int {

	return i.quantity

}

func (i *Invoice) SetPrice(price float64) {

	if price > 0.0 {

		i.price = price

	}

}

func (i *Invoice) GetPrice() float64 {

	return i.price

}

func (i *Invoice) GetInvoiceAmount() float64 {

	return i.price * float64(i.quantity)

}
