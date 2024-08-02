package builder

import "time"

type InvoiceBuilder struct {
	author string
	title string
	creationDate time.Time
	contentDraft string
	company *Company
}

func (builder *InvoiceBuilder) build() iTemplate {
	return &Invoice{
		author: builder.author,
		title: builder.title,
		creationDate: builder.creationDate,
		contentDraft: builder.contentDraft,
		company: builder.company,
	}
}

func (builder *InvoiceBuilder) withAuthor(author string)() {
	builder.author = author
}

func (builder *InvoiceBuilder) withTitle(title string)() {
	builder.title = title
}

func (builder *InvoiceBuilder) withCreationDate(creationDate time.Time) {
	builder.creationDate = creationDate
}

func (builder *InvoiceBuilder) withContentDraft(contentDraft string) {
	builder.contentDraft = contentDraft
}

func (builder *InvoiceBuilder) withCompany(name string, address string, representative string) {
	builder.company = &Company{
		name: name,
		address: address,
		representative: representative,
	}
}
