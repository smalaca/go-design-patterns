package command

type SuccessfulProductSaleProcess struct {
	product Product
}

func createSuccessfulProductSaleProcess(product Product) SuccessfulProductSaleProcess {
	return SuccessfulProductSaleProcess{
		product: product,
	}
}

func (controller *SuccessfulProductSaleProcess) informWarehouse(warehouseId int) {

}

func (controller *SuccessfulProductSaleProcess) notifyCustomer(customer Customer) {

}

func (controller *SuccessfulProductSaleProcess) createInvoice() {

}

func (controller *SuccessfulProductSaleProcess) execute() {

}
