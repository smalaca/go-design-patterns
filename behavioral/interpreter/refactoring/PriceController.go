package interpreter

type PriceController struct {
	isHolidaySeason bool
}

func createPriceController() PriceController {
	return PriceController{
		isHolidaySeason: false,
	}
}

func (controller *PriceController) priceFor(price int, customer Customer) int {
	if customer.wasLate {
		return price
	}

	newPrice := price

	if controller.isHolidaySeason {
		newPrice = newPrice - 10
	}

	if customer.isVip {
		newPrice = newPrice - 50
	}

	return newPrice
}

func (controller *PriceController) setHolidays(isHolidaySeason bool) {
	controller.isHolidaySeason = isHolidaySeason
}
