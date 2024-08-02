package abstractfactory

type LightThemeFactory struct {
}

func (factory *LightThemeFactory) createWindow() iWindow {
	return &LightWindow{}
}

func (factory *LightThemeFactory) createMenu() iMenu {
	return &LightMenu{}
}
