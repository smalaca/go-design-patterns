package iterator

type ComponentsGroup struct {
	name             string
	azureComponents  []Component
	awsComponents    []Component
	onPremComponents []Component
}

func (group *ComponentsGroup) iterator() Iterator {
	return createComponentsIterator(group.awsComponents, group.azureComponents, group.onPremComponents)
}
