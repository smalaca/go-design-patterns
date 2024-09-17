package chainofresponsibility

import "fmt"

type PaymentStep struct {
	product Product
	customer Customer
	next ProcessingStep
}

func (step *PaymentStep) execute() {
	fmt.Println("Payment for product: " + step.product.name + " by customer: " + step.customer.mail)

	// if (isSuccessfulPaymet()) {
	// 	step.successfulNext.execute()
	// } else {
	// 	step.failureNext.execute()
	// }

	step.next.execute()
}