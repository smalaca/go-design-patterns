package observer

import "fmt"

type Statistics struct {
}

func (observer *Statistics) notify(product Product) {
	fmt.Println("Stats about product: " + product.name)
}
