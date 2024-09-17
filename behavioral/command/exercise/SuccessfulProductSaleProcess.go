package command

type SuccessfulProductSaleProcess struct {
	commands []Command

}

func createSuccessfulProductSaleProcess() SuccessfulProductSaleProcess {
	return SuccessfulProductSaleProcess{
		commands: []Command{},
	}
}

func (controller *SuccessfulProductSaleProcess) informWarehouse(warehouseId int) {
	controller.commands = append(controller.commands, &InformWarehouseCommand{
		warehouseId: warehouseId,
	})
}

func (controller *SuccessfulProductSaleProcess) notifyCustomer(customer Customer) {
	controller.commands = append(controller.commands, &NotifyCustomerCommand{
		customer: customer,
	})
}

func (controller *SuccessfulProductSaleProcess) createInvoice() {
	controller.commands = append(controller.commands, &GenerateInvoiceCommand{})
}

func (controller *SuccessfulProductSaleProcess) execute(product Product) {
	for _, command := range controller.commands {
		command.execute(product)
	}
}
