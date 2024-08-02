package singleton

type DocumentTemplateController struct {
}

func (controller *DocumentTemplateController) execute(documentType string) {
	if documentType != "invoice" && documentType != "offer" {

	} else {
		// some other code
	}
}
