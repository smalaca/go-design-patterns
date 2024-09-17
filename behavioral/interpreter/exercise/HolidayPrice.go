package interpreter

type HolidayPrice struct {
	isHolidaySeason bool
	next            PriceExpression
}

func (expression *HolidayPrice) calculate(price int, customer Customer) int {
	if expression.isHolidaySeason {
		return expression.next.calculate(price-10, customer)
	}
	return expression.next.calculate(price, customer)
}
