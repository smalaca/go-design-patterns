package factormethod


type TranslatorController struct {
	language string
	germanDictionary GermanDictionary	
	googleTranslator GoogleTranslator
	spanishDictionary SpanishDictionary
	ancientSpanishDictionary AncientSpanishDictionary
}

func createTranslatorController() TranslatorController {
	return TranslatorController{
		germanDictionary: GermanDictionary{},
		googleTranslator: GoogleTranslator{},
		spanishDictionary: SpanishDictionary{},
		ancientSpanishDictionary: AncientSpanishDictionary{},
	}
}

func (controller *TranslatorController) change(language string) {
	controller.language = language
}

func (controller *TranslatorController) translate(content string) string {
	if (controller.language == "English") {
		return "Translation of :\"" + content + "\" in progress."
	} else if (controller.language == "Deutsch") {
		return "Übersetzung von:\"" + content + "\" in Arbeit."
	} else if (controller.language == "Spanish") {
		return "Traducción de:\"" + content + "\" en progreso."
	} else {
		return "I do not know the language"
	}
}
