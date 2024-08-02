package factormethod

type SpanishTranslatorFactory struct {

}

func (factory *SpanishTranslatorFactory) create() iTranslator {
	return &SpanishTranslator{
		newDictionary: SpanishDictionary{},
		oldDictionary: AncientSpanishDictionary{},
	}
}