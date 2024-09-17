package chainofresponsibility

type ProductSaleProcess struct {
	processingStep ProcessingStep
}

func createProductSaleProcess(customer Customer, product Product) ProductSaleProcess {
	invoiceStep := &InvoiceGenerationStep{
		product: product,
		customer: customer,
	}
	paymentStep := &PaymentStep{
		product: product,
		customer: customer,
		next: invoiceStep,
	}
	blockProductStep := &BlockProductStep{
		product: product,
		next: paymentStep,
	}

	return ProductSaleProcess{
		processingStep: blockProductStep,
	}
}

func (process *ProductSaleProcess) buy() {
	process.processingStep.execute()
}
