package chainofresponsibility

import "testing"

func Test_ChainOfResponsibility(t *testing.T) {
	product := Product{
		productId: 42,
		name:      "Design Patters Book",
	}
	customer := Customer{
		customerId: 987,
		mail:       "mailaddress@fake.com",
	}
	process := createProductSaleProcess(customer, product)

	process.buy()
}
