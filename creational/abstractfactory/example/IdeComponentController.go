package abstractfactory

type IdeComponentController struct {
	themeFactory iThemeFactory
}

func createIdeComponentController() IdeComponentController {
	return IdeComponentController{defaultThemeFactory()}
}

func (controller *IdeComponentController) changeTheme(themeName string) {
	if themeName == "DARK" {
		controller.themeFactory = &DarkThemeFactory{}
	} else if themeName == "LIGHT" {
		controller.themeFactory = &LightThemeFactory{}
	} else {
		controller.themeFactory = defaultThemeFactory()
	}
}

func defaultThemeFactory() iThemeFactory {
	return &DarkThemeFactory{}
}

func (controller *IdeComponentController) displayWindow(content string) {
	window := controller.themeFactory.createWindow()
	window.load(content)
	window.display()
}

func (controller *IdeComponentController) displayMenu(items []string) {
	window := controller.themeFactory.createMenu()
	window.load(items)
	window.display()
}
