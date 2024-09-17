package memento

import (
	"testing"
)

func Test_Memento(t *testing.T) {
	task := Task{
		taskId:             13,
		assigneeId:         42,
		name:               "Initial Name",
		description:        "Initial Description",
		acceptanceCriteria: "Initial Acceptance Criteria",
	}

	controller := createTaskController(task)
	controller.save()
	controller.displayTask()

	controller.changeName("Something name")
	controller.changeDescription("Another description")
	controller.changeAcceptanceCriteria("Some acceptance criteria")
	controller.save()
	controller.displayTask()

	controller.changeAcceptanceCriteria("Yet another change of acceptance criteria")
	controller.displayTask()

	controller.revertChange()
	controller.save()
	controller.displayTask()

	controller.revertChange()
	controller.revertChange()
	controller.displayTask()
}
