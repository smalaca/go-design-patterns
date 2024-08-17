package adapter

type TranslatorController struct {
	googleTranslator GoogleTranslator
	germanDictionary GermanDictionary
	language         string
}

func (controller *TranslatorController) change(language string) {
	controller.language = language
}

func (controller *TranslatorController) fromPolish(content string) string {
	switch controller.language {
	case "EN":
		request := Request{
			from:        "PL",
			to:          "EN",
			toTranslate: content,
		}
		response := controller.googleTranslator.translate(request)

		return response.translation
	case "DE":
		return controller.germanDictionary.translateFrom(content, "PL")
	}

	return "Unsupported"
}

func (controller *TranslatorController) toPolish(content string) string {
	switch controller.language {
		case "EN":
			request := Request{
				from:        "EN",
				to:          "PL",
				toTranslate: content,
			}
			response := controller.googleTranslator.translate(request)

			return response.translation
		case "DE":
			return controller.germanDictionary.translateTo(content, "PL")
	}

	return "Unsupported"
}

func (controller *TranslatorController) getGermanDictionary() GermanDictionary {
	if &controller.germanDictionary == nil {
		controller.germanDictionary = GermanDictionary{}
	}

	return controller.germanDictionary
}

func (controller *TranslatorController) getGoogleTranslator() GoogleTranslator {
	if &controller.googleTranslator == nil {
		controller.googleTranslator = GoogleTranslator{}
	}

	return controller.googleTranslator
}
