package proxy

import (
	"fmt"
	"testing"
)

func Test_Proxy(t *testing.T) {
	repository := ImageRepository{}
	image := repository.findByName("Hello World")
	
	fmt.Println(image.displayShort())
	fmt.Println(image.displayFull())
}
