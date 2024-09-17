package templatemethod

import "fmt"

type XmlFactory struct {
	ReportFactory
}

func createXmlFactory() XmlFactory {
	factory := XmlFactory{}
	factory.steps = &factory
	return factory
}

func (factory *XmlFactory) generateHeader(data ReportData) string {
	fmt.Println("XML: generate header")
	return "XML header"
}

func (factory *XmlFactory) generateContent(data ReportData) string {
	fmt.Println("XML: generate content")
	return "XML content"
}

func (factory *XmlFactory) generateFooter() string {
	fmt.Println("XML: generate footer")
	return "XML footer"
}