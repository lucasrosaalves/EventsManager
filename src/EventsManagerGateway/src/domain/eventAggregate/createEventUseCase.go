package eventAggregate

type ICreateEventUseCase interface {
	Execute(timestamp string, tag string, value string) bool
}
