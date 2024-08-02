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

func (offer *Offer) clone() iTemplate {
	return &Offer{
		author: offer.author,
		title: offer.title,
		creationDate: offer.creationDate,
		contentDraft: offer.contentDraft,
		company: offer.company,
		logo: offer.logo,
	}
}
