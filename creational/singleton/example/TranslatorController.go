package singleton

type TranslatorController struct {

}

func (controller *TranslatorController) execute(translatorType string) {
	if (translatorType != "english" && translatorType != "spanish" && translatorType != "deutsch") {
		getApplicationContext().incrementFailedOperations()
	} else {
		// some other code
	}
}