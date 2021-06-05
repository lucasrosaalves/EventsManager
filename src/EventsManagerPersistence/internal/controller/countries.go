package controller

import (
	"eventsmanagerpersistence/internal/usecases/query"
	"net/http"
)

type CountriesController struct {
	getCountriesQuery query.GetCountries
}

func NewCountriesController(getCountriesQuery query.GetCountries) *CountriesController {
	return &CountriesController{
		getCountriesQuery: getCountriesQuery,
	}
}

func (c *CountriesController) GetAll(writer http.ResponseWriter, r *http.Request) {

	countries := c.getCountriesQuery.Handle()

	Ok(writer, countries)
}
