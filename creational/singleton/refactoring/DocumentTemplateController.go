package singleton

type DocumentTemplateController struct {
	failuresStorage *FailuresStorage
}

func (controller *DocumentTemplateController) execute(documentType string) {
	if (documentType != "invoice" && documentType != "offer") {
		controller.failuresStorage.incrementFailedOperations()
	} else {
		// some other code
	}
}