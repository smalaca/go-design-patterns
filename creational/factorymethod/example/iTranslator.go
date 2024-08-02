package factormethod

type iTranslator interface {
	translate(content string) string
}