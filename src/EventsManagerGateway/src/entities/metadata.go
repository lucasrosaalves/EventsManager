package entities

import (
	"eventsmanagergateway/src/utils"
	"time"
)

type MetaData struct {
	CorrelationId string    `json:"correlation_id"`
	StartDate     time.Time `json:"start_date"`
}

func NewMetaData() *MetaData {
	return &MetaData{
		CorrelationId: utils.GenerateGuid(),
		StartDate:     time.Now(),
	}
}
