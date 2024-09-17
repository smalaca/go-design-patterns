package iterator

import (
	"fmt"
	"testing"
)

func Test_Iterator(t *testing.T) {
	controller := ComponentGroupController{}

	componentsGroup := controller.findOnAzureFor("Online Shop")

	iterator := componentsGroup.iterator()

	for iterator.hasNext() {
		componentName := iterator.getNext()
		fmt.Println(componentName)
	}
}
