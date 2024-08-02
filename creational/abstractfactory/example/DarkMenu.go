package abstractfactory

import (
	"fmt"
	"strconv"
)

type DarkMenu struct {
	items []string
}

func (menu *DarkMenu) display() {
	fmt.Println("Dark Menu:")
	for index, item := range menu.items {
		fmt.Println(strconv.Itoa(index + 1) + ": " + item)
	}
	fmt.Println()
}

func (menu *DarkMenu) load(items []string) {
	menu.items = items
}
