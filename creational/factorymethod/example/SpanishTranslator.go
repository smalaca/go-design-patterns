package factormethod

type SpanishTranslator struct {
	newDictionary SpanishDictionary
	oldDictionary AncientSpanishDictionary
}

func (translator *SpanishTranslator) translate(content string) string {
	return "Hola Mundo"
}