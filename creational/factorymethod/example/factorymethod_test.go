package factormethod

import (
	"fmt"
	"testing"
)

func Test_FactorMethod(t *testing.T) {
	controller := TranslatorController{}

	controller.change(&DeutschTranslatorFactory{})
	fmt.Println(controller.translate("Witaj Świecie"))

	controller.change(&EnglishTranslatorFactory{})
	fmt.Println(controller.translate("Witaj Świecie"))

	// controller.change("Spanish")
	// fmt.Println(controller.translate("Witaj Świecie"))
}
