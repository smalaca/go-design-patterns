package visitor

import "fmt"

type NotificationVisitor struct {

}

func (visitor *NotificationVisitor) visitTask(task *Task) {
	fmt.Println("Notification about task: " + task.name)
}

func (visitor *NotificationVisitor) visitStory(story *Story) {
	fmt.Println("Notification about story: " + story.name)
}

func (visitor *NotificationVisitor) visitProject(project *Project) {
	fmt.Println("No Notification about project")
}