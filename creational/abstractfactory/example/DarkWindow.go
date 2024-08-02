package abstractfactory

import "fmt"

type DarkWindow struct {
	content string
}

func (window *DarkWindow) display() {
	fmt.Print("Dark Window: " + window.content)
	fmt.Println()
}

func (window *DarkWindow) load(content string) {
	window.content = content
}