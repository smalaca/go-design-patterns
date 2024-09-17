package templatemethod

import (
	"testing"
)

func Test_TemplateMethod(t *testing.T) {
	data := ReportData{}
	pdfFactory := createPdfFactory()
	xmlFactory := createXmlFactory()

	pdfReport,_ := pdfFactory.create(data)
	pdfReport.display()
	xmlReport,_ := xmlFactory.create(data)
	xmlReport.display()
}
