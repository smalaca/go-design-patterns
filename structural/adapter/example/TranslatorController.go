package adapter

type TranslatorController struct {
	translator iTranslator
}

func (controller *TranslatorController) change(language string) {
	switch language {
		case "EN":
			controller.translator = &EnglishTranslator{GoogleTranslator{}}
		case "DE":
			controller.translator = &DeutschTranslator{GermanDictionary{}}
	}
}

func (controller *TranslatorController) fromPolish(content string) string {
	return "Translation from Polish: Input: \"" + content + "\"; Translation: \"" + controller.translator.fromPolish(content) + "\""
}

func (controller *TranslatorController) toPolish(content string) string {
	return "Translation to Polish: Input: \"" + content + "\"; Translation: \"" + controller.translator.toPolish(content) + "\""
}
