package chainofresponsibility

import "fmt"

type ProductSaleProcess struct {
	customer Customer
	product Product
}

func createProductSaleProcess(customer Customer, product Product) ProductSaleProcess {
	return ProductSaleProcess{
		customer: customer,
		product: product,
	}
}

func (process *ProductSaleProcess) buy() {
	fmt.Println("Blocking product: " + process.product.name)
	fmt.Println("Payment for product: " + process.product.name + " by customer: " + process.customer.mail)
	fmt.Println("Generating invoice for selling product: " + process.product.name + " for customer: " + process.customer.mail)
}
