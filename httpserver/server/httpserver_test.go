package server

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
}

func (s *StubPlayerStore) GetPlayerScore(player string) int {
	return s.scores[player]

}

func (s *StubPlayerStore) RecordWin(player string) {
	s.winCalls = append(s.winCalls, player)
}
func TestGETPlayers(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{
			"Pepper": 20,
			"Floyd":  10,
		},
		nil,
	}
	server := &PlayerServer{&store}
	t.Run("return Pepper's score", func(t *testing.T) {
		request := getScoreRequestByPlayerName("Pepper")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		// // PlayerServer(response, request)

		// got := response.Body.String()
		// want := "20"
		assertResponseStatusCode(t, response.Code, http.StatusOK)
		assertStringsEqual(t, response.Body.String(), "20")

	})

	t.Run("return Floyd's score", func(t *testing.T) {
		request := getScoreRequestByPlayerName("Floyd")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		// PlayerServer(response, request)

		// got := response.Body.String()
		// want := "10"
		assertResponseStatusCode(t, response.Code, http.StatusOK)
		assertStringsEqual(t, response.Body.String(), "10")

	})

	t.Run("returns 404 on missing players", func(t *testing.T) {
		request := getScoreRequestByPlayerName("Apollo")
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)
		assertResponseStatusCode(t, response.Code, http.StatusNotFound)
	})

	t.Run("returns accepted on POST", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodPost, "/players/Pepper", nil)
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)
		assertResponseStatusCode(t, response.Code, http.StatusAccepted)
	})

	t.Run("records win when POST", func(t *testing.T) {
		playerName := "Rafita"
		request := newPOSTWinRequest(playerName)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		assertResponseStatusCode(t, response.Code, http.StatusAccepted)
		if store.winCalls[len(store.winCalls)-1] != playerName {
			t.Errorf("wanted: %s but got: %s", playerName, store.winCalls[len(store.winCalls)-1])
		}
	})
}

func newPOSTWinRequest(player string) *http.Request {
	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", player), nil)
	return req
}

func assertStringsEqual(t *testing.T, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, wanted: %q", got, want)
	}
}

func getScoreRequestByPlayerName(player string) *http.Request {
	r, _ := http.NewRequest(http.MethodGet, "/players/"+player, nil)
	return r
}

func assertResponseStatusCode(t *testing.T, got int, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got status code: %d but wanted: %d", got, want)
	}
}

// func getPlayerScore(player string) string {
// 	switch player {
// 	case "Pepper":
// 		return "20"
// 	case "Floyd":
// 		return "10"
// 	default:
// 		return "0"
// 	}
// }
