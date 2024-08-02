package singleton

import (
	"fmt"
	"testing"
)

func Test_Singleton(t *testing.T) {
	failuresStorage := &FailuresStorage{}
	templateController := DocumentTemplateController{failuresStorage}
	translatorController := TranslatorController{failuresStorage}
	themeController := IdeComponentController{failuresStorage}

	fmt.Println(failuresStorage.getFailedOperations())

	templateController.execute("invoice")
	templateController.execute("wrong")
	templateController.execute("offer")
	templateController.execute("invalid")

	fmt.Println(failuresStorage.getFailedOperations())

	translatorController.execute("english")
	translatorController.execute("deutsch")
	translatorController.execute("deutsch")
	translatorController.execute("spanish")
	translatorController.execute("i do not know it")
	translatorController.execute("well, I should learn more")

	fmt.Println(failuresStorage.getFailedOperations())

	themeController.execute("light")
	themeController.execute("wrong")
	themeController.execute("wrong")
	themeController.execute("dark")
	themeController.execute("dark")
	themeController.execute("wrong")

	fmt.Println(failuresStorage.getFailedOperations())

	failuresStorage.reset()
	fmt.Println(failuresStorage.getFailedOperations())

	themeController.execute("light")
	themeController.execute("wrong")
	themeController.execute("wrong")
	themeController.execute("dark")
	themeController.execute("dark")
	themeController.execute("wrong")

	fmt.Println(failuresStorage.getFailedOperations())
}
