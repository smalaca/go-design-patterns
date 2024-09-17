package chainofresponsibility

import "fmt"

type BlockProductStep struct {
	product Product
	next ProcessingStep
}

func (step *BlockProductStep) execute() {
	fmt.Println("Blocking product: " + step.product.name)
	step.next.execute()
}
