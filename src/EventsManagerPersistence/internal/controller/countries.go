package controller

import (
	"encoding/json"
	"eventsmanagerpersistence/internal/entity"
	"net/http"
)

type CountriesController struct {
}

func NewCountriesController() *CountriesController {
	return &CountriesController{}
}

func (c *CountriesController) GetAll(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	content, _ := json.Marshal(&entity.Country{
		Id:   "1",
		Code: "BR",
		Name: "Brazil",
	})
	w.WriteHeader(200)
	w.Write(content)
}
