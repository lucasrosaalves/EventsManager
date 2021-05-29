package entities

import (
	"time"
)

type Event struct {
	Timestamp time.Time `json:"timestamp"`
	Country   string    `json:"country"`
	Region    string    `json:"region"`
	Sensor    string    `json:"sensor"`
	Value     string    `json:"value"`
	MetaData  *MetaData `json:"metaData"`
}
