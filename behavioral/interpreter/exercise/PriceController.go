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
	return price
}

func (controller *PriceController) setHolidays(isHolidaySeason bool) {
	controller.isHolidaySeason = isHolidaySeason
}
