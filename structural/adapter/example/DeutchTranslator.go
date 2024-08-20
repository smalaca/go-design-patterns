package adapter

type DeutchTranslator struct {
	dictionary GermanDictionary
}

func (translator *DeutchTranslator) fromPolish(content string) string {
	return translator.dictionary.translateFrom(content, "PL")
}

func (translator *DeutchTranslator) toPolish(content string) string {
	return translator.dictionary.translateTo(content, "PL")
}