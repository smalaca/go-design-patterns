package builder

import "time"

type OfferBuilder struct {
	author string
	title string
	creationDate time.Time
	contentDraft string
	company Company
}

func (builder *OfferBuilder) build() iTemplate {
	return &Offer{
		author: builder.author,
		title: builder.title,
		creationDate: builder.creationDate,
		contentDraft: builder.contentDraft,
		company: builder.company,
	}
}

func (builder *OfferBuilder) withAuthor(author string) {
	// validate if author exists
	builder.author = author
}

func (builder *OfferBuilder) withTitle(title string) {
	// validate if title contain more than 10 characters
	builder.title = title
}

func (builder *OfferBuilder) withCreationDate(creationDate time.Time) {
	// validate if creation date is from the past
	builder.creationDate = creationDate
}

func (builder *OfferBuilder) withContentDraft(contentDraft string) {
	builder.contentDraft = contentDraft
}

func (builder *OfferBuilder) withCompany(name string, address string, representative string) {
	builder.company = Company{
		name: name,
		address: address,
		representative: representative,
	}
}