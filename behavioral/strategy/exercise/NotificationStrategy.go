package strategy

type NotificationStrategy interface {
	send(message string)
}