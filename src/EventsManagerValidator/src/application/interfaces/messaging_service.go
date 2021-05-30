package interfaces

type MessagingService interface {
	Send(obj interface{}, queueName string) error
	Subscribe(queueName string, c chan []byte)
}
