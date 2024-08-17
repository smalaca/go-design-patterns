package facade

import (
	"testing"
)

func Test_Facade(t *testing.T) {
	smsGateway := SmsGateway{}
	mailClient := MailClient{} 
	chatService := ChatService{}

	title := "Design Patterns"
	message := "DP are great if you plan to keep your code mainteinable"


	smsGateway.send(SmsRequest{
		phoneNumber: "12345",
		message: message,
	})

	mailClient.sendMail(Mail{
		mailAddress: "sebastian.malaca@mymail.com",
		title: title,
		content: message,
	})

	chat := chatService.getFor(13)
	chat.send(message)
}
