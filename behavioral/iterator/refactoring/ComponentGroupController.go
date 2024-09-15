package iterator

type ComponentGroupController struct {
}

func (controller *ComponentGroupController) findOnAzureFor(name string) ComponentsGroup {
	azureComponents := []Component{
		Component{name: "notification service"},
		Component{name: "notification service"},
		Component{name: "accounting service"},
		Component{name: "invoice service"},
	}
	awsComponents := []Component{
		Component{name: "product catalogue"},
		Component{name: "product catalogue"},
		Component{name: "product catalogue"},
		Component{name: "product management"},
	}
	onPremComponents := []Component{
		Component{name: "product sales"},
		Component{name: "product sales"},
		Component{name: "payment service"},
	}

	return ComponentsGroup{
		name:             name,
		azureComponents:  azureComponents,
		awsComponents:    awsComponents,
		onPremComponents: onPremComponents,
	}
}
