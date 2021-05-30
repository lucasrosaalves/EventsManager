package domain

type CreateEventUseCase interface {
	Execute(timestamp int64, tag string, value string) error
}
