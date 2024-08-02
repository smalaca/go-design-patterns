package prototype

import (
	"fmt"
	"time"
)

type Invoice struct {
	author       string
	title        string
	creationDate time.Time
	contentDraft string
	company      *Company
	logo         *Logo
}

func (invoice *Invoice) print() {
	fmt.Println("-------------------------")
	fmt.Println("Invoice: " + invoice.author + "; " + invoice.title + "; " + invoice.contentDraft + "; " + invoice.creationDate.GoString() + "; ")
	invoice.logo.print()
	invoice.company.print()
	
	fmt.Println("-------------------------")
}

func (invoice *Invoice) clone() iTemplate {
	return &Invoice{
		author: invoice.author,
		title: invoice.title,
		creationDate: invoice.creationDate,
		contentDraft: invoice.contentDraft,
		company: invoice.company,
		logo: invoice.logo,
	}
}
