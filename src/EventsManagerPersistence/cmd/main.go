package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Result struct {
	Data string
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")

		content, _ := json.Marshal(&Result{
			Data: "OK",
		})
		w.WriteHeader(200)
		w.Write(content)

	}).Methods("GET")

	srv := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
