package singleton

import (
	"fmt"
	"testing"
)

func Test_Singleton(t *testing.T) {
	templateController := DocumentTemplateController{}
	translatorController := TranslatorController{}
	themeController := IdeComponentController{}

	fmt.Println(getApplicationContext().getFailedOperations())

	templateController.execute("invoice")
	templateController.execute("wrong")
	templateController.execute("offer")
	templateController.execute("invalid")

	fmt.Println(getApplicationContext().getFailedOperations())

	translatorController.execute("english")
	translatorController.execute("deutsch")
	translatorController.execute("deutsch")
	translatorController.execute("spanish")
	translatorController.execute("i do not know it")
	translatorController.execute("well, I should learn more")

	fmt.Println(getApplicationContext().getFailedOperations())

	themeController.execute("light")
	themeController.execute("wrong")
	themeController.execute("wrong")
	themeController.execute("dark")
	themeController.execute("dark")
	themeController.execute("wrong")

	fmt.Println(getApplicationContext().getFailedOperations())

	getApplicationContext().reset()
	fmt.Println(getApplicationContext().getFailedOperations())

	themeController.execute("light")
	themeController.execute("wrong")
	themeController.execute("wrong")
	themeController.execute("dark")
	themeController.execute("dark")
	themeController.execute("wrong")

	fmt.Println(getApplicationContext().getFailedOperations())
}
