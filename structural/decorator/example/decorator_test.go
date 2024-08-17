package decorator

import (
	"fmt"
	"testing"
)

func Test_Decorator(t *testing.T) {
	editor := TextEditor{}

	helloWorld := editor.write("Hello World!")
	helloWorld = editor.decorate(helloWorld, "B")
	helloWorld = editor.decorate(helloWorld, "I")
	helloWorld = editor.decorate(helloWorld, "U")

	fmt.Println(helloWorld.getValue())

	helloSebastian := editor.write("Hello Sebastian!")
	helloSebastian = editor.decorate(helloSebastian, "B")
	fmt.Println(helloSebastian.getValue())
}
