package visitor

type Visitor interface {
	visitTask(task *Task)
	visitStory(story *Story)
	visitProject(project *Project)
}