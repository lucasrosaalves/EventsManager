package server

import (
	"eventsmanagerpersistence/internal/controller"
	"eventsmanagerpersistence/internal/usecases/query"

	"github.com/gorilla/mux"
)

func configureRoutes() *mux.Router {
	controller := controller.NewCountriesController(query.NewGetCountries())

	router := mux.NewRouter()
	router.HandleFunc("/", controller.GetAll).Methods("GET")
	return router
}
