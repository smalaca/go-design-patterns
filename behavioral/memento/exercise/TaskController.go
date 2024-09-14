package memento

import "fmt"

type TaskController struct {
	task Task
}

func createTaskController(task Task) TaskController {
	return TaskController{
		task: task,
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

}

func (controller *TaskController) revertChange() {

}

// func (controller *TaskController) popLastSnapshot(snapshots *[]TaskMemento) TaskMemento {
// 	f := len(*snapshots)
// 	lastSnapshot := (*snapshots)[f-1]
// 	*snapshots = (*snapshots)[:f-1]

// 	return lastSnapshot
// }
