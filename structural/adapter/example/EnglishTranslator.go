package adapter

type EnglishTranslator struct {
	translator GoogleTranslator
}

func (translator *EnglishTranslator) fromPolish(content string) string {
	request := Request{
		from: "PL",
		to: "EN",
		toTranslate: content,
	}
	response := translator.translator.translate(request)

	return response.translation
}

func (translator *EnglishTranslator) toPolish(content string) string {
	request := Request{
		from: "EN",
		to: "PL",
		toTranslate: content,
	}
	response := translator.translator.translate(request)
	
	return response.translation
}
