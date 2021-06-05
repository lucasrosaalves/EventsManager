package server

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Server struct {
}

func New() *Server {
	return &Server{}
}

func (*Server) Run() {
	router := mux.NewRouter()

	for _, r := range ConfigureRoutes() {
		router.HandleFunc(r.Path, r.Handler).Methods(r.Method)
	}

	srv := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Print("Running at 8000")
	log.Fatal(srv.ListenAndServe())
}
