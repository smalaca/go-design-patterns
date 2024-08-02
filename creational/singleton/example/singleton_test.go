package singleton

import (
	"testing"
)

func Test_Singleton(t *testing.T) {
	templateController := DocumentTemplateController{}
	translatorController := TranslatorController{}
	themeController := IdeComponentController{}

	templateController.execute("invoice")
	templateController.execute("wrong")
	templateController.execute("offer")
	templateController.execute("invalid")

	translatorController.execute("english")
	translatorController.execute("deutsch")
	translatorController.execute("deutsch")
	translatorController.execute("spanish")
	translatorController.execute("i do not know it")
	translatorController.execute("well, I should learn more")

	themeController.execute("light")
	themeController.execute("wrong")
	themeController.execute("wrong")
	themeController.execute("dark")
	themeController.execute("dark")
	themeController.execute("wrong")

	themeController.execute("light")
	themeController.execute("wrong")
	themeController.execute("wrong")
	themeController.execute("dark")
	themeController.execute("dark")
	themeController.execute("wrong")
}
