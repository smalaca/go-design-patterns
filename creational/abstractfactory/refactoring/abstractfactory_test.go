package abstractfactory

import (
	"testing"
)

func Test_AbstractFactory(t *testing.T) {
	items := []string{"File", "Edit", "Selection", "View"}
	ide := createIdeComponentController()

	ide.changeTheme("DARK")
	ide.displayWindow("creational patterns are the best!")
	ide.displayMenu(items)

	ide.changeTheme("LIGHT")
	ide.displayWindow("things are great when you know them")
	ide.displayMenu(items)

	ide.changeTheme("UNDEFINED SO SHOULD BE DEFAULT")
	ide.displayWindow("creational patterns are event better!")
	ide.displayMenu(items)
}
