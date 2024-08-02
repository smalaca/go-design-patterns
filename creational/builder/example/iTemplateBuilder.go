package builder

import "time"

type iTemplateBuilder interface {
	build() iTemplate

	withAuthor(author string)
	withTitle(title string)
	withCreationDate(creationDate time.Time)
	withContentDraft(contentDraft string)
	withCompany(name string, address string, representative string)
}