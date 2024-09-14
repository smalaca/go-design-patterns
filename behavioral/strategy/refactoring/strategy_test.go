package strategy

import (
	"testing"
)

func Test_Strategy(t *testing.T) {
	controller := createNotificationController()
	message := "Very important message"

	controller.choose("MAIL")
	controller.send(message)

	controller.choose("SMS")
	controller.send(message)

	controller.choose("CHAT")
	controller.send(message)
}
