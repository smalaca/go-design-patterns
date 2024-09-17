package templatemethod

import "fmt"

type PdfFactory struct {
	ReportFactory
}

func createPdfFactory() PdfFactory {
	factory := PdfFactory{}
	factory.steps = &factory
	return factory
}

func (factory *PdfFactory) generateHeader(data ReportData) string {
	fmt.Println("PDF: generate header")
	return "PDF header"
}

func (factory *PdfFactory) generateContent(data ReportData) string {
	fmt.Println("PDF: generate content")
	return "PDF content"
}

func (factory *PdfFactory) generateFooter() string {
	fmt.Println("PDF: generate footer")
	return "PDF footer"
}