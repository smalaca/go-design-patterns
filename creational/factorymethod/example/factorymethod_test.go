package factormethod

import (
	"fmt"
	"testing"
)

func Test_FactorMethod(t *testing.T) {
	controller := TranslatorController{}

	controller.change("Deutsch")
	fmt.Println(controller.translate("Hello World"))

	controller.change("English")
	fmt.Println(controller.translate("Hello World"))

	controller.change("Spanish")
	fmt.Println(controller.translate("Hello World"))
}
