package builder

import (
	"testing"
	"time"
)

func Test_Builder(t *testing.T) {
	controller := DocumentTemplateController{}

	offerTemplate := controller.createTemplate(dtoWithCompany("offer"))
	display(offerTemplate)

	invoiceTemplate := controller.createTemplate(dtoWithCompany("invoice"))
	display(invoiceTemplate)
    
	offerTemplateWithoutCompany := controller.createTemplate(dtoWithoutCompany("offer"))
	display(offerTemplateWithoutCompany)
    
	invoiceTemplateWithoutCompany := controller.createTemplate(dtoWithoutCompany("invoice"))
	display(invoiceTemplateWithoutCompany)
}

func display(template iTemplate) {
	template.print()
}

func dtoWithCompany(documentType string) TemplateDto {
    return TemplateDto{
        documentType: documentType,
        author: "Sebastian Malaca",
        title: "QWE/2024/8/13",
        creationDate: time.Now(),
        contentDraft: "anything or everything or nothing",
        companyName: "Let's Talk About Quality",
        companyAddress: "existing one",
        companyRepresentative: "Sebastian Malaca",
    }
}

func dtoWithoutCompany(documentType string) TemplateDto {
    return TemplateDto{
        documentType: documentType,
        author: "Sebastian Malaca",
        title: "QWE/2024/8/13",
        creationDate: time.Now(),
        contentDraft: "anything or everything or nothing",
    }
}
