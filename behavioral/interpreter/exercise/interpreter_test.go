package interpreter

import (
	"fmt"
	"strconv"
	"testing"
)

func Test_Interpreter(t *testing.T) {
	controller := createPriceController()

	controller.setHolidays(true)
	price := controller.priceFor(1000, Customer{
		isVip:   false,
		wasLate: false,
	})
	fmt.Println("Price during holidays is: " + strconv.Itoa(price))
	price = controller.priceFor(1000, Customer{
		isVip:   true,
		wasLate: false,
	})
	fmt.Println("Price during holidays is: " + strconv.Itoa(price))

	price = controller.priceFor(1000, Customer{
		isVip:   true,
		wasLate: true,
	})
	fmt.Println("Price during holidays is: " + strconv.Itoa(price))

	controller.setHolidays(false)
	price = controller.priceFor(1000, Customer{
		isVip:   false,
		wasLate: false,
	})
	fmt.Println("Price after holidays is: " + strconv.Itoa(price))
	price = controller.priceFor(1000, Customer{
		isVip:   true,
		wasLate: false,
	})
	fmt.Println("Price after holidays is: " + strconv.Itoa(price))

	price = controller.priceFor(1000, Customer{
		isVip:   true,
		wasLate: true,
	})
	fmt.Println("Price after holidays is: " + strconv.Itoa(price))
}
