package domain

import (
	"errors"
	"strings"
	"time"
)

const (
	queueName = "event.received"
)

type Event struct {
	Timestamp   time.Time `json:"timestamp"`
	CountryName string    `json:"country_name"`
	RegionName  string    `json:"region_name"`
	SensorName  string    `json:"sensor_name"`
	Value       string    `json:"value"`
	MetaData    *MetaData `json:"metaData"`
}

func NewEvent(timestamp int64, tag string, value string, correlationId string) (*Event, error) {
	country, region, sensor, tagError := handleTag(tag)

	if tagError != nil {
		return nil, tagError
	}

	return &Event{
		Timestamp:   time.Unix(timestamp, 0),
		CountryName: country,
		RegionName:  region,
		SensorName:  sensor,
		Value:       value,
		MetaData:    NewMetaData(correlationId),
	}, nil
}

func (*Event) GetQueueName() string {
	return queueName
}

func handleTag(tag string) (string, string, string, error) {
	tagSplited := strings.Split(tag, ".")

	if tagSplited == nil || len(tagSplited) != 3 {
		return "", "", "", errors.New("Invalid tag")
	}

	return tagSplited[0], tagSplited[1], tagSplited[2], nil
}
