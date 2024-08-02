package abstractfactory

import (
	"fmt"
	"strconv"
)

type LightMenu struct {
	items []string
}

func (menu *LightMenu) display() {
	fmt.Println("Light Menu:")
	for index, item := range menu.items {
		fmt.Println(strconv.Itoa(index + 1) + ": " + item)
	}
	fmt.Println()
}

func (menu *LightMenu) load(items []string) {
	menu.items = items
}
