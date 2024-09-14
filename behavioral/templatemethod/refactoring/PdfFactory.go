package templatemethod

import (
	"errors"
	"fmt"
)

type PdfFactory struct {

}

func createPdfFactory() PdfFactory {
	return PdfFactory{}
}

func (factory *PdfFactory) create(data ReportData) (*Report, error) {
	if factory.areValid(data) {
		return &Report{
			header: factory.createHeader(data),
			content: factory.createContent(data),
			footer: "PDF footer",
			signature: "signature",
		}, nil
	} else {
		return nil, errors.New("No file found in this path")
	}
}

func (factory *PdfFactory) areValid(data ReportData) bool {
	return true
}

func (factory *PdfFactory) createHeader(data ReportData) string {
	fmt.Println("PDF: generate header")
	return "PDF header"
}

func (factory *PdfFactory) createContent(data ReportData) string {
	fmt.Println("PDF: generate content")
	return "PDF content"
}