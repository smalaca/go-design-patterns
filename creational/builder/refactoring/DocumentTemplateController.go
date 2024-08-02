package builder

import "strings"

type DocumentTemplateController struct {
}

func (controller *DocumentTemplateController) createTemplate(dto TemplateDto) iTemplate {
	switch dto.documentType {
	case "offer":
		return controller.newOffer(dto)
	case "invoice":
		return controller.newInvoice(dto)
	default:
		return controller.newOffer(dto)
	}
}

func (controller *DocumentTemplateController) newOffer(dto TemplateDto) *Offer {
	offer := &Offer{
		author: dto.author,
		title: dto.title,
		creationDate: dto.creationDate,
		contentDraft: dto.contentDraft,
	}

	if (len(strings.Trim(dto.companyName, " ")) > 0) {
		company := Company{
			name: dto.companyName,
			address: dto.companyAddress,
			representative: dto.companyRepresentative,
		}
		offer.company = &company
	}

	return offer
}

func (controller *DocumentTemplateController) newInvoice(dto TemplateDto) *Invoice {
	invoice := &Invoice{
		author: dto.author,
		title: dto.title,
		creationDate: dto.creationDate,
		contentDraft: dto.contentDraft,
	}

	if (len(strings.Trim(dto.companyName, " ")) > 0) {
		company := Company{
			name: dto.companyName,
			address: dto.companyAddress,
			representative: dto.companyRepresentative,
		}
		invoice.company = &company
	}

	return invoice
}
