package prototype

import "fmt"

type Company struct {
	name           string
	address        string
	representative string
}

func (company *Company) print() {
	fmt.Println("Company: " + company.name + "; " + company.address + "; " + company.representative)
}
