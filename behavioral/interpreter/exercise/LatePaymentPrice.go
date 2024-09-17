package interpreter

type LatePaymentPrice struct {
	originalPrice int
	next          PriceExpression
}

func (expression *LatePaymentPrice) calculate(price int, customer Customer) int {
	if customer.wasLate {
		return expression.originalPrice
	}
	return expression.next.calculate(price, customer)
}
