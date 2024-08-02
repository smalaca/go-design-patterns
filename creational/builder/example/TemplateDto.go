package builder

import "time"

type TemplateDto struct {
	documentType string
	author string
	title string
	creationDate time.Time
	contentDraft string
	companyName string
	companyAddress string
	companyRepresentative string
}