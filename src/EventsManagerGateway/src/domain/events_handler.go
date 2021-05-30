package domain

type EventsHandler interface {
	Execute(timestamp int64, tag string, value string) error
}
