package adapter

type GermanDictionary struct {
}

func (translator *GermanDictionary) translateFrom(content string, language string) string {
	return "Hallo Welt"
}

func (translator *GermanDictionary) translateTo(content string, language string) string {
	return "Witaj Świecie"
}