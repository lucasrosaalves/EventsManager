package server

import (
	"eventsmanagerpersistence/internal/controller"
	"net/http"
)

type Route struct {
	Path    string
	Method  string
	Handler func(http.ResponseWriter, *http.Request)
}

func NewRoute(path string, method string, handler func(http.ResponseWriter, *http.Request)) *Route {
	return &Route{
		Path:    path,
		Method:  method,
		Handler: handler,
	}
}

func ConfigureRoutes() []Route {
	controller := controller.NewCountriesController()
	return []Route{
		{Path: "/", Method: "GET", Handler: controller.GetAll},
	}
}
