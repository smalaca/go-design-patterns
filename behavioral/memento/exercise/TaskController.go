package memento

import "fmt"

type TaskController struct {
	task      Task
	snapshots []TaskMemento
}

func createTaskController(task Task) TaskController {
	return TaskController{
		task:      task,
		snapshots: []TaskMemento{},
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
	controller.snapshots = append(controller.snapshots, controller.task.createMemento())
}

func (controller *TaskController) revertChange() {
	snapshot := controller.popLastSnapshot(&controller.snapshots)
	controller.task.apply(snapshot)
}

func (controller *TaskController) popLastSnapshot(snapshots *[]TaskMemento) TaskMemento {
	f := len(*snapshots)
	lastSnapshot := (*snapshots)[f-1]
	*snapshots = (*snapshots)[:f-1]

	return lastSnapshot
}
