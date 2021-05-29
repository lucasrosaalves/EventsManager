package entities

import (
	"time"
)

type MetaData struct {
	CorrelationId string    `json:"correlation_id"`
	StartDate     time.Time `json:"start_date"`
}
