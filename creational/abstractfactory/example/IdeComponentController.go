package abstractfactory

type IdeComponentController struct {
	
}

func createIdeComponentController() IdeComponentController {
	return IdeComponentController{}
}

func (controller *IdeComponentController) changeTheme(themeName string) {

}

func (controller *IdeComponentController) displayWindow(content string) {
	
}

func (controller *IdeComponentController) displayMenu(items []string) {
	
}
