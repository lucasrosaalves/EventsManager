package domain

import (
	"strings"
	"time"
)

const (
	queueName = "event.validated"
)

type EventValidated struct {
	Timestamp time.Time `json:"timestamp"`
	CountryId string    `json:"country_id"`
	RegionId  string    `json:"region_id"`
	SensorId  string    `json:"sensor_id"`
	Value     string    `json:"value"`
	Status    string    `json:"status"`
	MetaData  MetaData  `json:"metaData"`
}

func NewEventValidated(timestamp time.Time, countryId string, regionId string, sensorId string, value string, metadata MetaData) *EventValidated {

	return &EventValidated{
		Timestamp: timestamp,
		CountryId: countryId,
		RegionId:  regionId,
		SensorId:  sensorId,
		Value:     value,
		Status:    getStatus(value),
		MetaData:  metadata,
	}
}

func (*EventValidated) GetQueueName() string {
	return queueName
}

func getStatus(value string) string {
	if strings.TrimSpace(value) == "" {
		return "Error"
	}

	return "Processed"
}
