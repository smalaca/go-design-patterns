package strategy

import "fmt"

type NotificationController struct {
	choice string
}

func createNotificationController() NotificationController {
	return NotificationController{choice: "SMS"}
}

func (controller *NotificationController) choose(choice string) {
	controller.choice = choice
}

func (controller *NotificationController) send(message string) {
	if (controller.choice == "MAIL") {
		controller.sendMail(message)
	} else if (controller.choice == "SMS") {
		controller.sendSms(message)
	} else if (controller.choice == "CHAT") {
		controller.sendChat(message)
	}
}

func (controller *NotificationController) sendMail(message string) {
	fmt.Println("MAIL: " + message)
}

func (controller *NotificationController) sendSms(message string) {
	fmt.Println("SMS: " + message)
}

func (controller *NotificationController) sendChat(message string) {
	fmt.Println("CHAT: " + message)
}