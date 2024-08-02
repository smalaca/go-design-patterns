package prototype

import (
	"fmt"
	"time"
)

type Offer struct {
	author       string
	title        string
	creationDate time.Time
	contentDraft string
	company      *Company
	logo         *Logo
}

func (offer *Offer) print() {
	fmt.Println("-------------------------")
	fmt.Println("Offer: " + offer.author + "; " + offer.title + "; " + offer.contentDraft + "; " + offer.creationDate.GoString() + "; ")
	offer.logo.print()
	offer.company.print()
	fmt.Println("-------------------------")
}
