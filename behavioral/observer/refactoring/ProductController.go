package observer

import "strconv"

type ProductController struct {
	customerService *CustomerService
	statistics      *Statistics
	warehouse       *Warehouse
}

func createProductController() ProductController {
	return ProductController{}
}

func (controller *ProductController) buy(productId int) {
	if controller.exists(productId) {
		product := controller.findById(productId)
		product.bought()
		controller.notifyAll(product)
	}
}

func (controller *ProductController) exists(productId int) bool {
	return true
}

func (controller *ProductController) findById(productId int) Product {
	return Product{
		productId: productId,
		name:      "Product No: " + strconv.Itoa(productId),
	}
}

func (controller *ProductController) notifyAll(product Product) {
	if (controller.customerService != nil) {
		controller.customerService.notify(product)
	}
	if (controller.warehouse != nil) {
		controller.warehouse.notify(product)
	}
	if (controller.statistics != nil) {
		controller.statistics.notify(product)
	}
}

func (controller *ProductController) registerCustomerService(customerService CustomerService) {
	controller.customerService = &customerService
}

func (controller *ProductController) registerStatistics(statistics Statistics) {
	controller.statistics = &statistics
}

func (controller *ProductController) registerWarehouse(warehouse Warehouse) {
	controller.warehouse = &warehouse
}

func (controller *ProductController) unregisterCustomerService() {
	controller.customerService = nil
}

func (controller *ProductController) unregisterStatistics() {
	controller.statistics = nil
}

func (controller *ProductController) unregisterWarehouse() {
	controller.warehouse = nil
}
