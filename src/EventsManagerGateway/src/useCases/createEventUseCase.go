package useCases

type CreateEventUseCase struct {
}

func (*CreateEventUseCase) Execute(timestamp string, tag string, value string) bool {
	return true
}
