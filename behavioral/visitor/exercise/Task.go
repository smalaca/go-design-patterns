package visitor

type Task struct {
	taskId int
	storyId int
	assigneeId int
	name string
}

func (task *Task) accept(visitor Visitor) {
	visitor.visitTask(task)
}
