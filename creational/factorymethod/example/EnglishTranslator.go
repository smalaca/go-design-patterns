package factormethod

type EnglishTranslator struct {
	googleTranslator GoogleTranslator
}

func (translator *EnglishTranslator) translate(content string) string {
	return "Hello World"
}