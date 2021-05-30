package domain

type EventsHandler interface {
	Execute(event EventReceived) error
}
