package query

import "eventsmanagerpersistence/internal/entity"

type GetCountries interface {
	Handle() *[]entity.Country
}

type getCountries struct {
	repository entity.CountryRepository
}

func NewGetCountries(repository entity.CountryRepository) GetCountries {
	return &getCountries{
		repository: repository,
	}
}

func (query *getCountries) Handle() *[]entity.Country {
	return query.repository.GetAll()
}
