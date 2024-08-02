package abstractfactory

import "fmt"

type LightWindow struct {
	content string
}

func (window *LightWindow) display() {
	fmt.Print("Light Window: " + window.content)
	fmt.Println()
}

func (window *LightWindow) load(content string) {
	window.content = content
}
