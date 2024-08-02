package builder

import "strings"


type DocumentTemplateController struct {
}

func (controller *DocumentTemplateController) createTemplate(dto TemplateDto) iTemplate {
	builder := controller.createBuilder(dto.documentType)
	builder.withAuthor(dto.author)
	builder.withTitle(dto.title)
	builder.withCreationDate(dto.creationDate)
	builder.withContentDraft(dto.contentDraft)

	if (len(strings.Trim(dto.companyName, " ")) > 0) {
		builder.withCompany(dto.companyName, dto.companyAddress, dto.companyRepresentative)
	}

	return builder.build()
}

func (controller *DocumentTemplateController) createBuilder(documentType string) iTemplateBuilder {
	if (documentType == "offer") {
		return &OfferBuilder{}
	} else {
		return nil
	}
}
