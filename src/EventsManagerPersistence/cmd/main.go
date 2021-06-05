package main

import (
	"eventsmanagerpersistence/internal/server"
)

func main() {
	server := server.New()

	server.Run()
}
