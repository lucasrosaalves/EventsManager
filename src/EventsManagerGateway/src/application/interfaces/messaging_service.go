package interfaces

type MessagingService interface {
	Send(obj interface{}, queueName string) error
}
