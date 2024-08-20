package adapter

import (
	"fmt"
	"testing"
)

func Test_Adapter(t *testing.T) {
	controller := TranslatorController{}

	controller.change("DE")
	fmt.Println(controller.toPolish("Hallo Welt"))
	fmt.Println(controller.fromPolish("Witaj Świecie"))

	// controller.change("EN")
	// fmt.Println(controller.toPolish("Hello World"))
	// fmt.Println(controller.fromPolish("Witaj Świecie"))
}
