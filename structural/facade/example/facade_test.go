package facade

import (
	"testing"
)

func Test_Facade(t *testing.T) {
	facade := NotificationFacade{}
	user := User{
		id: 13, 
		phoneNumber: "12345",
		mailAddress: "sebastian.malaca@mymail.com",
	}
	message := Message{
		title: "Design Patterns",
		content: "DP are great if you plan to keep your code mainteinable",
	}


	facade.sendSms(user, message)
	facade.sendMail(user, message)
	facade.sendChatMessage(user, message)
}
