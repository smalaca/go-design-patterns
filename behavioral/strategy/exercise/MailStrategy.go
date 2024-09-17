package strategy

import "fmt"

type MailStrategy struct {
	
}

func (strategy *MailStrategy) send(message string) {
	fmt.Println("MAIL: " + message)
}