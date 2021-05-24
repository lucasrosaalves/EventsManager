package eventAggregate

type Event struct {
	Timestamp string `json:"timestamp" binding:"required"`
	Tag       string `json:"tag" binding:"required"`
	Value     string `json:"value"`
}
