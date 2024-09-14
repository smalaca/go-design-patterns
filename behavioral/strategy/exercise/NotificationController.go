package strategy

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

}
