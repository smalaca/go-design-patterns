package command

import "testing"

func Test_Command(t *testing.T) {
	product := Product{
		productId: 42,
		name: "Design Patters Book",
	}
	controller := createSuccessfulProductSaleProcess()

	// w momencie dodawania produktu, różne dla produktów
	controller.informWarehouse(13)
	controller.notifyCustomer(Customer{
		customerId: 987,
		mail: "mailaddress@fake.com",
	})
	controller.createInvoice()

	// po zakupie produktu
	controller.execute(product)
}
