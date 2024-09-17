package chainofresponsibility

import "fmt"

type InvoiceGenerationStep struct {
	product Product
	customer Customer
}

func (step *InvoiceGenerationStep) execute() {
	fmt.Println("Generating invoice for selling product: " + step.product.name + " for customer: " + step.customer.mail)
}