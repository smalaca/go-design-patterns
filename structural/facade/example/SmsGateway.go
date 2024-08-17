package facade

import "fmt"

type SmsRequest struct{
	phoneNumber string
	message string
}

type SmsGateway struct {

}

func (gateway *SmsGateway) send(request SmsRequest) {
	fmt.Println("Sending SMS to: " + request.phoneNumber + "; message: " + request.message)
}