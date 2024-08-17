package facade

import "fmt"

type Chat struct {
	userId int
}

type ChatService struct {

}

func (service *ChatService) getFor(userId int) Chat {
	return Chat{userId}
}

func (chat *Chat) send(message string) {
	fmt.Println("Sending Chat Message to:", chat.userId, "; message:", message)
}