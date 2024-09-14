package observer

type Product struct {
	productId int
	name string
	isBought bool
}

func (product *Product) bought() {
	product.isBought = true
}