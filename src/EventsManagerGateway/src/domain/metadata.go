package domain

import (
	"time"
)

type MetaData struct {
	CorrelationId string    `json:"correlation_id"`
	CreatedAt     time.Time `json:"created_at"`
}

func NewMetaData(correlationId string) *MetaData {
	return &MetaData{
		CorrelationId: correlationId,
		CreatedAt:     time.Now(),
	}
}
