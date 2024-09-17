package command

import "fmt"

type NotifyCustomerCommand struct {
	customer Customer
}

func (command *NotifyCustomerCommand) execute(product Product) {
	fmt.Println("Customer: " + command.customer.mail + " was informed about:" + product.name)
}