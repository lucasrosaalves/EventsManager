package interfaces

type IEventsService interface {
	Execute(timestamp string, tag string, value string) bool
}
