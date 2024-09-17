package iterator

type ComponentsIterator struct {
	currentIndex     int
	azureComponents  []Component
	awsComponents    []Component
	onPremComponents []Component
}

func createComponentsIterator(awsComponents []Component, azureComponents []Component, onPremComponents []Component) Iterator {
	return &ComponentsIterator{
		currentIndex:     0,
		awsComponents:    awsComponents,
		azureComponents:  azureComponents,
		onPremComponents: onPremComponents,
	}
}

func (iterator *ComponentsIterator) hasNext() bool {
	return (iterator.currentIndex + 1) <= iterator.numberOfAllComponents()
}

func (iterator *ComponentsIterator) numberOfAllComponents() int {
	return len(iterator.azureComponents) + len(iterator.awsComponents) + len(iterator.onPremComponents)
}

func (iterator *ComponentsIterator) getNext() string {
	componentName := ""
	if (iterator.currentIndex < len(iterator.azureComponents)) {
		componentName = iterator.azureComponents[iterator.currentIndex].name
	} else if (iterator.currentIndex - len(iterator.azureComponents) < len(iterator.awsComponents)) {
		componentName = iterator.awsComponents[iterator.currentIndex - len(iterator.azureComponents)].name
	} else {
		componentName = iterator.onPremComponents[iterator.currentIndex - len(iterator.azureComponents) - len(iterator.awsComponents)].name
	}

	iterator.currentIndex++
	return componentName
}
