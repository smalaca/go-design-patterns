package abstractfactory

import (
	"fmt"
	"strconv"
)

type IdeComponentController struct {
	themeName string
}

func createIdeComponentController() IdeComponentController {
	return IdeComponentController{
		themeName: "DARK",
	}
}

func (controller *IdeComponentController) changeTheme(themeName string) {
	controller.themeName = themeName
}

func (controller *IdeComponentController) displayWindow(content string) {
	if controller.themeName == "DARK" {
		fmt.Print("Dark Window: " + content)
		fmt.Println()
	} else if controller.themeName == "LIGHT" {
		fmt.Print("Light Window: " + content)
		fmt.Println()
	} else {
		fmt.Print("Dark Window: " + content)
		fmt.Println()
	}
}

func (controller *IdeComponentController) displayMenu(items []string) {
	if controller.themeName == "DARK" {
		fmt.Println("Dark Menu:")
		for index, item := range items {
			fmt.Println(strconv.Itoa(index + 1) + ": " + item)
		}
		fmt.Println()
	} else if controller.themeName == "LIGHT" {
		fmt.Println("Light Menu:")
		for index, item := range items {
			fmt.Println(strconv.Itoa(index + 1) + ": " + item)
		}
		fmt.Println()
	} else {
		fmt.Println("Dark Menu:")
		for index, item := range items {
			fmt.Println(strconv.Itoa(index + 1) + ": " + item)
		}
		fmt.Println()
	}
}
