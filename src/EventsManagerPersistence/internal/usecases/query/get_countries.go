package query

import "eventsmanagerpersistence/internal/entity"

type GetCountries interface {
	Handle() *[]entity.Country
}

type getCountries struct {
}

func NewGetCountries() GetCountries {
	return &getCountries{}
}

func (query *getCountries) Handle() *[]entity.Country {
	return &[]entity.Country{
		{
			Id:   "1",
			Code: "BR",
			Name: "BR",
		},
		{
			Id:   "2",
			Code: "USA",
			Name: "USA",
		},
	}
}
