package templatemethod

import (
	"errors"
	"fmt"
)

type ReportSteps interface {
	generateHeader(data ReportData) string
	generateContent(data ReportData) string
	generateFooter() string
}

type ReportFactory struct {
	steps ReportSteps
}

func (factory *ReportFactory) create(data ReportData) (*Report, error) {
	if factory.areValid(data) {

		
		return &Report{
			header: factory.steps.generateHeader(data),
			content: factory.steps.generateContent(data),
			footer: factory.steps.generateFooter(),
			signature: factory.generateSignature(),
		}, nil
	} else {
		return nil, errors.New("No file found in this path")
	}
}

func (factory *ReportFactory) areValid(data ReportData) bool {
	fmt.Println("TEMPLATE METHOD: validate parameters")
	return true
}

func (factory *ReportFactory) generateSignature() string {
	fmt.Println("TEMPLATE METHOD: generate signature")
	return "TEMPLATE METHOD: signature"
}
