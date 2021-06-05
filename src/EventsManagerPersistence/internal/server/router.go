package server

import (
	"eventsmanagerpersistence/internal/controller"
	"eventsmanagerpersistence/internal/infra/repository"
	"eventsmanagerpersistence/internal/usecases/query"

	"github.com/gorilla/mux"
)

func configureRoutes() *mux.Router {
	controller := controller.NewCountriesController(query.NewGetCountries(
		repository.NewCountryRepository(),
	))

	router := mux.NewRouter()
	router.HandleFunc("/", controller.GetAll).Methods("GET")
	return router
}
