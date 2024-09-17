package interpreter

type RawPrice struct {
}

func (expression *RawPrice) calculate(price int, customer Customer) int {
	return price
}
