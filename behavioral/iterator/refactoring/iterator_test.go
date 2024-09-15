package iterator

import (
	"fmt"
	"testing"
)

func Test_Iterator(t *testing.T) {
	controller := ComponentGroupController{}

	componentsGroup := controller.findOnAzureFor("Online Shop")

	currentIndex := 0
	componentName := ""

	for (currentIndex + 1) <= len(componentsGroup.azureComponents)+len(componentsGroup.awsComponents)+len(componentsGroup.onPremComponents) {
		if currentIndex < len(componentsGroup.azureComponents) {
			componentName = componentsGroup.azureComponents[currentIndex].name
		} else if currentIndex-len(componentsGroup.azureComponents) < len(componentsGroup.awsComponents) {
			componentName = componentsGroup.awsComponents[currentIndex-len(componentsGroup.azureComponents)].name
		} else {
			componentName = componentsGroup.onPremComponents[currentIndex-len(componentsGroup.azureComponents)-len(componentsGroup.awsComponents)].name
		}

		fmt.Println(componentName)
		currentIndex++
	}
}
