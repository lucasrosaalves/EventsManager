package useCases

type CreateEventUseCase struct {
}

type ICreateEventUseCase interface {
	Execute(timestamp string, tag string, value string) bool
}

func (*CreateEventUseCase) Execute(timestamp string, tag string, value string) bool {
	return true
}
