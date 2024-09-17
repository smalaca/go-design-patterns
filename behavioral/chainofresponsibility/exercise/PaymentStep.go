package chainofresponsibility

import "fmt"

type PaymentStep struct {
	product  Product
	customer Customer
	next     ProcessStep
}

func (step *PaymentStep) execute() {
	fmt.Println("Payment for product: " + step.product.name + " by customer: " + step.customer.mail)
	step.next.execute()
}
