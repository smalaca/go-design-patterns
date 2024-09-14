package templatemethod

type PdfFactory struct {
	
}

func createPdfFactory() PdfFactory {
	return PdfFactory{}
}

func (factory *PdfFactory) create(data ReportData) (*Report, error) {
	return &Report{}, nil
}