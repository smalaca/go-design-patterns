package templatemethod

import (
	"errors"
	"fmt"
)

type XmlFactory struct {

}

func createXmlFactory() XmlFactory {
	return XmlFactory{}
}

func (factory *XmlFactory) create(data ReportData) (*Report, error) {
	if factory.areValid(data) {		
		return &Report{
			header: factory.generateHeader(data),
			content: factory.generateContent(data),
			footer: factory.generateFooter(),
			signature: factory.generateSignature(),
		}, nil
	} else {
		return nil, errors.New("No file found in this path")
	}
}

func (factory *XmlFactory) areValid(data ReportData) bool {
	return true
}

func (factory *XmlFactory) generateSignature() string {
	return "signature"
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