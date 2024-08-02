package factormethod

type DeutschTranslatorFactory struct {
	
}

func (factory *DeutschTranslatorFactory) create() iTranslator {
	dictionary := GermanDictionary{}
	// connect with dictionary
	// validate connectivity

	return &DeutschTranslator{dictionary: dictionary}
}