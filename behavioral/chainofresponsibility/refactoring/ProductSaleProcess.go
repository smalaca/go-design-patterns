package chainofresponsibility

import "fmt"

type ProductSaleProcess struct {
	product  Product
	customer Customer
}

func createProductSaleProcess(customer Customer, product Product) ProductSaleProcess {
	return ProductSaleProcess{
		product:  product,
		customer: customer,
	}
}

func (process *ProductSaleProcess) buy() {
	fmt.Println("Blocking product: " + process.product.name)
	fmt.Println("Payment for product: " + process.product.name + " by customer: " + process.customer.mail)
	fmt.Println("Generating invoice for selling product: " + process.product.name + " for customer: " + process.customer.mail)
}
