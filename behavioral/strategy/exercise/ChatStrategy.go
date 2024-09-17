package strategy

import "fmt"

type ChatStrategy struct {
	
}

func (strategy *ChatStrategy) send(message string) {
	fmt.Println("CHAT: " + message)
}