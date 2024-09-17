package visitor

type Story struct {
	storyId int
	projectId int
	sprintId int
	name string
}

func (story *Story) accept(visitor Visitor) {
	visitor.visitStory(story)
}
