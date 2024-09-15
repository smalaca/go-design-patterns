package command

import (
	"fmt"
	"strconv"
)

type SuccessfulProductSaleProcess struct {
	product Product
	shouldInformWarehouse bool
	shouldInformCustomer bool
	shouldGenerateInvoice bool
	warehouseId int
	customer Customer
}

func createSuccessfulProductSaleProcess(product Product) SuccessfulProductSaleProcess {
	return SuccessfulProductSaleProcess{
		product: product,
		shouldInformWarehouse: false,
		shouldInformCustomer: false,
		shouldGenerateInvoice: false,
	}
}

func (controller *SuccessfulProductSaleProcess) informWarehouse(warehouseId int) {
	controller.warehouseId = warehouseId
	controller.shouldInformWarehouse = true
}

func (controller *SuccessfulProductSaleProcess) notifyCustomer(customer Customer) {
	controller.customer = customer
	controller.shouldInformCustomer = true
}

func (controller *SuccessfulProductSaleProcess) createInvoice() {
	controller.shouldGenerateInvoice = true
}

func (controller *SuccessfulProductSaleProcess) execute() {
	if controller.shouldInformWarehouse {
		fmt.Println("Sending information to warehouse: " + strconv.Itoa(controller.warehouseId) + ", about product: " + controller.product.name)
	}

	if controller.shouldInformCustomer {
		fmt.Println("Sending information to customer: " + controller.customer.mail + ", about product: " + controller.product.name)
	}

	if controller.shouldGenerateInvoice {
		fmt.Println("Generating invoice for selling product: " + controller.product.name)
	}
}
