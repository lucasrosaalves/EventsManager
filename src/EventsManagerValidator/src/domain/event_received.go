package domain

import (
	"time"
)

type EventReceived struct {
	Timestamp   time.Time `json:"timestamp"`
	CountryName string    `json:"country_name"`
	RegionName  string    `json:"region_name"`
	SensorName  string    `json:"sensor_name"`
	Value       string    `json:"value"`
	MetaData    MetaData  `json:"metaData"`
}
