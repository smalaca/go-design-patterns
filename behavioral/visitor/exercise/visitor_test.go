package visitor

import (
	"testing"
)

func Test_Visitor(t *testing.T) {
	controller := IssueController{}
	task := &Task{
		taskId: 13,
		storyId: 42,
		assigneeId: 123,
		name: "The Task",
	}
	story := &Story{
		storyId: 42,
		projectId: 987,
		sprintId: 645,
		name: "Yet Another Story",
	}
	project := &Project{
		projectId: 987,
		stakeholder: "Captain America",
		name: "Avengers",
	}

	notificationVisitor := &NotificationVisitor{}
	statisticsVisitor := &StatisticsVisitor{}

	controller.register(notificationVisitor)
	controller.processFor(task)
	controller.processFor(story)
	controller.processFor(project)

	controller.register(statisticsVisitor)
	controller.processFor(task)
	controller.processFor(story)
	controller.processFor(project)
}
