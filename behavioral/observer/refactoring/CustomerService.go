package observer

import "fmt"

type CustomerService struct {
}

func (observer *CustomerService) notify(product Product) {
	fmt.Println("Customer informed about product: " + product.name)
}
