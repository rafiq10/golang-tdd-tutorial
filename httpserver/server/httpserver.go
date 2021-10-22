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

type PlayerStore interface {
	GetPlayerScore(player string) int
}

type PlayerServer struct {
	Store PlayerStore
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")
	score := p.Store.GetPlayerScore(player)
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
