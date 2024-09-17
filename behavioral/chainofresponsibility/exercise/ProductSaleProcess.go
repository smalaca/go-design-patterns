package chainofresponsibility

type ProductSaleProcess struct {
	step ProcessStep
}

func createProductSaleProcess(customer Customer, product Product) ProductSaleProcess {
	return ProductSaleProcess{
		step: &BlockProductStep{
			product: product,
			next: &PaymentStep{
				product:  product,
				customer: customer,
				next: &IvoiceGenerationStep{
					product:  product,
					customer: customer,
				},
			},
		},
	}
}

func (process *ProductSaleProcess) buy() {
	process.step.execute()
}
