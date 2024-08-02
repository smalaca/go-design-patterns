package singleton

type IdeComponentController struct {

}

func (controller *IdeComponentController) execute(theme string) {
	if (theme != "light" && theme != "dark") {
		getApplicationContext().incrementFailedOperations()
	} else {
		// some other code
	}
}