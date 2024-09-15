package chainofresponsibility

type ProductSaleProcess struct {
	
}

func createProductSaleProcess(customer Customer, product Product) ProductSaleProcess {
	return ProductSaleProcess{}
}

func (process *ProductSaleProcess) buy() {
	
}
