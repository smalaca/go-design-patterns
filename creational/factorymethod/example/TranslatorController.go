package factormethod

type TranslatorController struct {
	translator iTranslator
}

func (controller *TranslatorController) change(factory iTranslatorFactory) {
	controller.translator = factory.create()
}

func (controller *TranslatorController) translate(content string) string {
	return controller.translator.translate(content)
}
