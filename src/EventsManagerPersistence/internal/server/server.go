package server

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type Server struct {
}

func New() *Server {
	return &Server{}
}

func (*Server) Run() {
	router := configureRoutes()

	srv := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Print("Running at 8000")
	log.Fatal(srv.ListenAndServe())
}
