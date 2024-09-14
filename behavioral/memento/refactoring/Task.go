package memento

import "strconv"

type Task struct {
	taskId             int
	assigneeId         int
	name               string
	description        string
	acceptanceCriteria string
}

func (task *Task) summary() string {
	return "Task[" +
		"taskId: " + strconv.Itoa(task.taskId) + ";" +
		"assigneeId: " + strconv.Itoa(task.assigneeId) + ";" +
		"name: " + task.name + ";" +
		"description: " + task.description + ";" +
		"acceptanceCriteria: " + task.acceptanceCriteria + ";" +
		"]"
}

func (task *Task) changeName(name string) {
	task.name = name
}

func (task *Task) changeDescription(description string) {
	task.description = description
}

func (task *Task) changeAcceptanceCriteria(acceptanceCriteria string) {
	task.acceptanceCriteria = acceptanceCriteria
}
