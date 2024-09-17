package chainofresponsibility

import "fmt"

type IvoiceGenerationStep struct {
	product Product
	customer Customer
}

func (step *IvoiceGenerationStep) execute() {
	fmt.Println("Generating invoice for selling product: " + step.product.name + " for customer: " + step.customer.mail)
}
