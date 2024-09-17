package strategy

import "fmt"

type SmsStrategy struct {
	
}

func (strategy *SmsStrategy) send(message string) {
	fmt.Println("SMS: " + message)
}