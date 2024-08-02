package builder

import (
	"fmt"
	"time"
)

type Offer struct {
	author string
	title string
	creationDate time.Time
	contentDraft string
	company *Company
}

func (offer *Offer) print() {
	fmt.Println("-------------------------")
	fmt.Println("Offer: " + offer.author + "; " + offer.title + "; " + offer.contentDraft + "; " + offer.creationDate.GoString() + "; ")
	if offer.company != nil {
		offer.company.print()
	}
	fmt.Println("-------------------------")
}