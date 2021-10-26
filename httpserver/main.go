package main

import (
	"log"
	"net/http"

	"tdd-tutorial.com/httpserver/server"
)

func main() {
	store := server.NewInMemoryPlayerStore()
	server := &server.PlayerServer{store}
	// handler := http.HandlerFunc(server.PlayerServer.Ser)
	log.Fatal(http.ListenAndServe(":5000", server))
}
