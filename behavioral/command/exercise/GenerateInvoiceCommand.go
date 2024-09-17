package command

import "fmt"

type GenerateInvoiceCommand struct {

}

func (command *GenerateInvoiceCommand) execute(product Product) {
	fmt.Println("Invoice genered for:" + product.name)
}