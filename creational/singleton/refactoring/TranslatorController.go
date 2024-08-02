package singleton

type TranslatorController struct {
	failuresStorage *FailuresStorage
}

func (controller *TranslatorController) execute(translatorType string) {
	if (translatorType != "english" && translatorType != "spanish" && translatorType != "deutsch") {
		controller.failuresStorage.incrementFailedOperations()
	} else {
		// some other code
	}
}