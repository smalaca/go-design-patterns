package observer

import "fmt"

type Warehouse struct {
}

func (observer *Warehouse) notify(product Product) {
	fmt.Println("Warehose informed about product: " + product.name)
}
