package abstractfactory

type DarkThemeFactory struct {
}

func (factory *DarkThemeFactory) createWindow() iWindow {
	return &DarkWindow{}
}

func (factory *DarkThemeFactory) createMenu() iMenu {
	return &DarkMenu{}
}
