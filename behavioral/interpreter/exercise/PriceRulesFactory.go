package interpreter

type PriceRulesFactory struct {
}

func (factory *PriceRulesFactory) raw() PriceExpression {
	return &RawPrice{}
}

func (factory *PriceRulesFactory) latePayment(originalPrice int, expression PriceExpression) PriceExpression {
	return &LatePaymentPrice{
		originalPrice: originalPrice,
		next:          expression,
	}
}

func (factory *PriceRulesFactory) vipPrice(expression PriceExpression) PriceExpression {
	return &VipPrice{next: expression}
}

func (factory *PriceRulesFactory) holiday(isHolidaySeason bool, expression PriceExpression) PriceExpression {
	return &HolidayPrice{
		isHolidaySeason: isHolidaySeason,
		next:            expression,
	}
}
