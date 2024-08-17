package facade

import "fmt"

type Mail struct {
	mailAddress string
	title string
	content string
}

type MailClient struct {

}

func (client *MailClient) sendMail(mail Mail) {
	fmt.Println("Sending Mail to: " + mail.mailAddress + "; title: " + mail.title + "; content: " + mail.content)
}