package prototype

import (
	"fmt"
	"testing"
)

func Test_Prototype(t *testing.T) {
	controller := createDocumentTemplateController()

	offerTemplate := controller.createTemplate("offer")
	display(offerTemplate)

	invoiceTemplate := controller.createTemplate("invoice")
	display(invoiceTemplate)
}

func display(template string) {
	fmt.Println(template)
}
