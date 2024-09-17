package interpreter

type VipPrice struct {
	next PriceExpression
}

func (expression *VipPrice) calculate(price int, customer Customer) int {
	if customer.isVip {
		return expression.next.calculate(price-50, customer)
	}
	return expression.next.calculate(price, customer)
}
