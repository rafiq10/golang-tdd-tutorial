package main

import (
	"log"
	"net/http"

	"tdd-tutorial.com/httpserver/server"
)

type InMemoryStore struct {
	score int
}

func (i *InMemoryStore) GetPlayerScore(player string) int {
	return 123
}

func main() {
	store := &InMemoryStore{}
	server := &server.PlayerServer{store}
	// handler := http.HandlerFunc(server.PlayerServer.Ser)
	log.Fatal(http.ListenAndServe(":5000", server))
}
