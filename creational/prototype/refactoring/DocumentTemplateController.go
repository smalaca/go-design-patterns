package prototype

import "time"

type DocumentTemplateController struct {
}

func createDocumentTemplateController() DocumentTemplateController {
	return DocumentTemplateController{}
}

func (controller *DocumentTemplateController) createTemplate(templateType string) iTemplate {
	if (templateType == "offer") {
		return &Offer{
			author:       "Sebastian Malaca",
			title:        "QWE/2024/8/13",
			contentDraft: "anything or everything or nothing",
			creationDate: time.Now(),
			company:      controller.createCompany(),
			logo:         &Logo{},
		}
	} else {
		return &Invoice{
			author:       "Sebastian Malaca",
			title:        "QWE/2024/8/13",
			contentDraft: "anything or everything or nothing",
			creationDate: time.Now(),
			company:      controller.createCompany(),
			logo:         &Logo{},
		}
	}
}

func (controller *DocumentTemplateController) createCompany() *Company {
	return &Company{
		name:           "Let's Talk About Quality",
		address:        "existing one",
		representative: "Sebastian Malaca",
	}
}