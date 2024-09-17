package visitor

type IssueController struct {
	visitor Visitor
}

func (controller *IssueController) processFor(issue Issue) {
	issue.accept(controller.visitor)
}

func (controller *IssueController) register(visitor Visitor) {
	controller.visitor = visitor
}
