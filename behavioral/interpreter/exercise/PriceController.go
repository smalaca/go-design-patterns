package interpreter

type PriceController struct {
	isHolidaySeason   bool
	priceRulesFactory PriceRulesFactory
}

func createPriceController() PriceController {
	return PriceController{
		isHolidaySeason:   false,
		priceRulesFactory: PriceRulesFactory{},
	}
}

func (controller *PriceController) priceFor(price int, customer Customer) int {
	expression := controller.priceRulesFactory.raw()
	expression = controller.priceRulesFactory.latePayment(price, expression)
	expression = controller.priceRulesFactory.vipPrice(expression)
	expression = controller.priceRulesFactory.holiday(controller.isHolidaySeason, expression)

	return expression.calculate(price, customer)
}

func (controller *PriceController) setHolidays(isHolidaySeason bool) {
	controller.isHolidaySeason = isHolidaySeason
}
