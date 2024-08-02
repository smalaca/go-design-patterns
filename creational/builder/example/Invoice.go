package builder

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
}

func (invoice *Invoice) print() {
	fmt.Println("-------------------------")
	fmt.Println("Invoice: " + invoice.author + "; " + invoice.title + "; " + invoice.contentDraft + "; " + invoice.creationDate.GoString() + "; ")
	if (invoice.company != nil) {
		invoice.company.print()
	}
	fmt.Println("-------------------------")
}
