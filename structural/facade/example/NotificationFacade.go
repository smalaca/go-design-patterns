package facade

type NotificationFacade struct {
	smsGateway SmsGateway
	mailClient MailClient
	chatService ChatService
}

type User struct {
	id int
	phoneNumber string
	mailAddress string
}

type Message struct {
	title string
	content string
}

func (facade *NotificationFacade) sendSms(user User, message Message) {
	request := SmsRequest{
		phoneNumber: user.phoneNumber,
		message: message.content,
	}
	facade.smsGateway.send(request)
}

func (facade *NotificationFacade) sendMail(user User, message Message) {
	mail := Mail{
		mailAddress: user.mailAddress,
		title: message.title,
		content: message.content,
	}
	facade.mailClient.sendMail(mail)
}

func (facade *NotificationFacade) sendChatMessage(user User, message Message) {
	chat := facade.chatService.getFor(user.id)
	chat.send(message.content)
}