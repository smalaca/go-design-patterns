package observer

import (
	"testing"
)

func Test_Observer(t *testing.T) {
	controller := createProductController()
	statistics := Statistics{}
	customerService := CustomerService{}
	warehouse := Warehouse{}

	controller.registerStatistics(statistics)
	controller.registerCustomerService(customerService)
	controller.registerWarehouse(warehouse)

	controller.buy(13)
	controller.buy(42)

	controller.unregisterCustomerService()

	controller.buy(7)
}
