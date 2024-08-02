package abstractfactory

type iThemeFactory interface {
	createWindow() iWindow
	createMenu() iMenu
}
