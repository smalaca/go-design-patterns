package singleton

type IdeComponentController struct {
	failuresStorage *FailuresStorage
}

func (controller *IdeComponentController) execute(theme string) {
	if (theme != "light" && theme != "dark") {
		controller.failuresStorage.incrementFailedOperations()
	} else {
		// some other code
	}
}