package factormethod

type EnglishTranslatorFactory struct {

}

func (factory *EnglishTranslatorFactory) create() iTranslator {
	return &EnglishTranslator{GoogleTranslator{}}
}