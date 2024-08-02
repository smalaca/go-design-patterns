package prototype

import "time"

type DocumentTemplateController struct {
	prototypes map[string]iTemplate
}

func createDocumentTemplateController() DocumentTemplateController {
	company := &Company{
		name:           "Let's Talk About Quality",
		address:        "existing one",
		representative: "Sebastian Malaca",
	}
	offer := &Offer{
		author:       "Sebastian Malaca",
		title:        "QWE/2024/8/13",
		contentDraft: "anything or everything or nothing",
		creationDate: time.Now(),
		company:      company,
		logo:         &Logo{},
	}

	invoice := &Invoice{
		author:       "Sebastian Malaca",
		title:        "QWE/2024/8/13",
		contentDraft: "anything or everything or nothing",
		creationDate: time.Now(),
		company:      company,
		logo:         &Logo{},
	}

	prototypes := map[string]iTemplate{
		"invoice": invoice,
		"offer": offer,
	}

	return DocumentTemplateController{prototypes: prototypes}
}

func (controller *DocumentTemplateController) createTemplate(templateType string) iTemplate {
	return controller.prototypes[templateType].clone()
}
