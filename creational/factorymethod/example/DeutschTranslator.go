package factormethod

type DeutschTranslator struct {
	dictionary GermanDictionary
}

func (translator *DeutschTranslator) translate(content string) string {
	return "Hallo Welt";
}