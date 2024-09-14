package visitor

import "fmt"

type NotificationService struct {

}

func (visitor *NotificationService) notifyAboutTask(task *Task) {
	fmt.Println("Notification about task: " + task.name)
}

func (visitor *NotificationService) notifyAboutStory(story *Story) {
	fmt.Println("Notification about story: " + story.name)
}