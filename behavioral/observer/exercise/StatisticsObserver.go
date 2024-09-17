package observer

import "fmt"

type StatisticsObserver struct {

}

func (observer *StatisticsObserver) notify(product Product) {
	fmt.Println("Stats about product: " + product.name)
}

func (observer *StatisticsObserver) name() string {
	return "statistics"
}