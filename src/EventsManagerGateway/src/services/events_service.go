package services

type EventsService struct {
}

func (*EventsService) Execute(timestamp string, tag string, value string) bool {
	return true
}
