package observer

import "fmt"

type CustomerServiceObserver struct {
}

func (observer *CustomerServiceObserver) notify(product Product) {
	fmt.Println("Customer informed about product: " + product.name)
}

func (observer *CustomerServiceObserver) name() string {
	return "customerservice"
}
