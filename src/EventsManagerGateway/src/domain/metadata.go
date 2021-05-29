package domain

import (
	"time"
)

type MetaData struct {
	CorrelationId string    `json:"correlation_id"`
	StartDate     time.Time `json:"start_date"`
}

func NewMetaData(correlationId string) *MetaData {
	return &MetaData{
		CorrelationId: correlationId,
		StartDate:     time.Now(),
	}
}
