package adapter

type iTranslator interface {
	fromPolish(content string) string
	toPolish(content string) string
}
