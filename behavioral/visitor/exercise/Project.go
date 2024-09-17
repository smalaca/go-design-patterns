package visitor

type Project struct {
	projectId   int
	name        string
	stakeholder string
}

func (project *Project) accept(visitor Visitor) {
	visitor.visitProject(project)
}
