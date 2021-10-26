package server

import (
	"fmt"
	"net/http"
	"strings"
)

// func PlayerServer(w http.ResponseWriter, r *http.Request) {
// 	player := strings.TrimPrefix(r.URL.Path, "/players/")

// 	score := getPlayerScore(player)
// 	PrintStringoResponseWriter(w, score)
// }

// type InMemoryStore struct {
// 	score int
// }

// func (i *InMemoryStore) GetPlayerScore(player string) int {
// 	return 123
// }
// func (i *InMemoryStore) RecordWin(player string) {

// }

type PlayerStore interface {
	GetPlayerScore(player string) int
	RecordWin(player string)
}

type PlayerServer struct {
	Store PlayerStore
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	playerName := strings.TrimPrefix(r.URL.Path, "/players/")
	switch r.Method {
	case http.MethodPost:
		p.processWin(w, playerName)
	case http.MethodGet:
		p.showScore(w, playerName)
	}

}

func (p *PlayerServer) processWin(w http.ResponseWriter, playerName string) {
	// playerName := strings.TrimPrefix(r.URL.Path, "/players/")
	p.Store.RecordWin(playerName)
	w.WriteHeader(http.StatusAccepted)
}
func (p *PlayerServer) showScore(w http.ResponseWriter, playerName string) {
	score := p.Store.GetPlayerScore(playerName)
	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.WriteHeader(http.StatusOK)
	}
	PrintIntToResponseWriter(w, score)
}
func PrintErrorToResponseWriter(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprintf(w, "invalid request: "+err.Error())
}

func PrintIntToResponseWriter(w http.ResponseWriter, out int) {
	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprintf(w, fmt.Sprint(out))
}
