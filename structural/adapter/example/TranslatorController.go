package adapter

type TranslatorController struct {
	translator iTranslator
}

func (controller *TranslatorController) change(language string) {
	
}

func (controller *TranslatorController) fromPolish(content string) string {
	return "Translation from Polish: Input: \"" + content + "\"; Translation: \"" + "NOT SUPPORTED NOW" + "\""
}

func (controller *TranslatorController) toPolish(content string) string {
	return "Translation to Polish: Input: \"" + content + "\"; Translation: \"" + "NOT SUPPORTED NOW" + "\""
}
