package strategy

type NotificationController struct {
	services map[string]NotificationStrategy
	choice string
}

func createNotificationController() NotificationController {
	services := map[string]NotificationStrategy{
		"MAIL": &MailStrategy{},
		"SMS":  &SmsStrategy{},
		"CHAT": &ChatStrategy{},
	}

	return NotificationController{services: services, choice: "SMS"}
}

func (controller *NotificationController) choose(choice string) {
	controller.choice = choice
}

func (controller *NotificationController) send(message string) {
	controller.services[controller.choice].send(message)
}