package command

import "fmt"

type InformWarehouseCommand struct {
	warehouseId int
}

func (command *InformWarehouseCommand) execute(product Product) {
	fmt.Println("Warehouse was informed about:" + product.name)
}