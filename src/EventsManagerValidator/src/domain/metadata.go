package domain

import (
	"time"
)

type MetaData struct {
	CorrelationId string    `json:"correlation_id"`
	CreatedAt     time.Time `json:"created_at"`
}
