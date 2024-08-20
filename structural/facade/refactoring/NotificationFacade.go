package facade


type Message struct {
	userId int
	phoneNumber string
	mailAddress string
	title string
	message string
}

type NotificationFacade struct {
	smsGateway SmsGateway
	mailClient MailClient
	chatService ChatService
}

func (facade *NotificationFacade) sendSms(message Message) {
	facade.smsGateway.send(SmsRequest{
		phoneNumber: message.phoneNumber,
		message: message.message,
	})
}

func (facade *NotificationFacade) sendMail(message Message) {
	facade.mailClient.sendMail(Mail{
		mailAddress: message.mailAddress,
		title: message.title,
		content: message.message,
	})
}

func (facade *NotificationFacade) sendChatMessage(message Message) {
	chat := facade.chatService.getFor(message.userId)
	chat.send(message.message)
}