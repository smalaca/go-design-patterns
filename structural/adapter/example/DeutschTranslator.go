package adapter

type DeutschTranslator struct {
	dictionary GermanDictionary
}

func (translator *DeutschTranslator) fromPolish(content string) string {
	return translator.dictionary.translateFrom(content, "PL")
}

func (translator *DeutschTranslator) toPolish(content string) string {
	return translator.dictionary.translateTo(content, "PL")
}
