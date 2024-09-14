package memento

import "fmt"

type TaskController struct {
	task                Task
	names               []string
	descriptions        []string
	acceptanceCriterias []string
}

func createTaskController(task Task) TaskController {
	return TaskController{
		task:                task,
		names:               []string{},
		descriptions:        []string{},
		acceptanceCriterias: []string{},
	}
}

func (controller *TaskController) displayTask() {
	fmt.Println(controller.task.summary())
}

func (controller *TaskController) changeName(name string) {
	controller.task.changeName(name)
}

func (controller *TaskController) changeDescription(description string) {
	controller.task.changeDescription(description)
}

func (controller *TaskController) changeAcceptanceCriteria(acceptanceCriteria string) {
	controller.task.changeAcceptanceCriteria(acceptanceCriteria)
}

func (controller *TaskController) save() {
	controller.names = append(controller.names, controller.task.name)
	controller.descriptions = append(controller.descriptions, controller.task.description)
	controller.acceptanceCriterias = append(controller.acceptanceCriterias, controller.task.acceptanceCriteria)
}

func (controller *TaskController) revertChange() {
	name := controller.popLastSnapshot(&controller.names)
	description := controller.popLastSnapshot(&controller.descriptions)
	acceptanceCriteria := controller.popLastSnapshot(&controller.acceptanceCriterias)
	controller.task.changeName(name)
	controller.task.changeDescription(description)
	controller.task.changeAcceptanceCriteria(acceptanceCriteria)
}

func (controller *TaskController) popLastSnapshot(snapshots *[]string) string {
	f := len(*snapshots)
	lastSnapshot := (*snapshots)[f-1]
	*snapshots = (*snapshots)[:f-1]

	return lastSnapshot
}
