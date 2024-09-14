package templatemethod

import "fmt"

type Report struct {
	header string
	content string
	footer string
	signature string
}

func (report *Report) display() {
	fmt.Println("Report: [" +
		"Header: " + report.header + "; " +
		"Content: " + report.content + "; " +
		"Footer: " + report.footer + "; " +
		"Signature: " + report.signature + "; " +
		"]")
}