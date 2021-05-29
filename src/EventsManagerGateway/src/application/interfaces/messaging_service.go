package interfaces

type MessagingService interface {
	Send(obj interface{}) error
}
