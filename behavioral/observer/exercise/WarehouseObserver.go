package observer

import "fmt"

type WarehouseObserver struct {

}

func (observer *WarehouseObserver) notify(product Product) {
	fmt.Println("Warehose informed about product: " + product.name)
}

func (observer *WarehouseObserver) name() string {
	return "warehouse"
}