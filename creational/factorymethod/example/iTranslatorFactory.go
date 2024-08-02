package factormethod

type iTranslatorFactory interface {
	create() iTranslator
}