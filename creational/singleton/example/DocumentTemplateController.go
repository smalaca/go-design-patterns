package singleton

type DocumentTemplateController struct {

}

func (controller *DocumentTemplateController) execute(documentType string) {
	if (documentType != "invoice" && documentType != "offer") {
		getApplicationContext().incrementFailedOperations()
	} else {
		// some other code
	}
}