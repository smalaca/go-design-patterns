package observer

import (
	"testing"
)

func Test_Observer(t *testing.T) {
	controller := createProductController()
	statisticsObserver := StatisticsObserver{}
	customerServiceObserver := CustomerServiceObserver{}
	warehouseObserver := WarehouseObserver{}

	controller.register(&statisticsObserver)
	controller.register(&customerServiceObserver)
	controller.register(&warehouseObserver)

	controller.buy(13)
	controller.buy(42)

	controller.unregister(&customerServiceObserver)

	controller.buy(7)
}
