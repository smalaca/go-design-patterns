package facade

import (
	"testing"
)

func Test_Facade(t *testing.T) {
	notification := NotificationFacade{}

	title := "Design Patterns"
	message := "DP are great if you plan to keep your code mainteinable"

	notification.sendSms(Message{
		phoneNumber: "12345",
		message: message,
	})

	notification.sendMail(Message{
		mailAddress: "sebastian.malaca@mymail.com",
		title: title,
		message: message,
	})

	notification.sendChatMessage(Message{
		userId: 13,
		message: message,
	})
}
