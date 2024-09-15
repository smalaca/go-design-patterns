package iterator

type ComponentsGroup struct {
	name             string
	azureComponents  []Component
	awsComponents    []Component
	onPremComponents []Component
}
