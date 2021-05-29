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
	Timestamp time.Time `json:"timestamp"`
	Country   string    `json:"country"`
	Region    string    `json:"region"`
	Sensor    string    `json:"sensor"`
	Value     string    `json:"value"`
	MetaData  *MetaData `json:"metaData"`
}

func NewEvent(timestamp int64, tag string, value string, correlationId string) (*Event, error) {
	country, region, sensor, tagError := handleTag(tag)

	if tagError != nil {
		return nil, tagError
	}

	return &Event{
		Timestamp: time.Unix(timestamp, 0),
		Country:   country,
		Region:    region,
		Sensor:    sensor,
		Value:     value,
		MetaData:  NewMetaData(correlationId),
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
