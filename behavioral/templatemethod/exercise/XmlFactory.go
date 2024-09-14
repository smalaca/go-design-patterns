package templatemethod

type XmlFactory struct {
}

func createXmlFactory() XmlFactory {
	return XmlFactory{}
}

func (factory *XmlFactory) create(data ReportData) (*Report, error) {
	return &Report{}, nil
}
