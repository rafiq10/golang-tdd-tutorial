package server

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type StubPlayerStore struct {
	scores map[string]int
}

func (s *StubPlayerStore) GetPlayerScore(player string) int {
	return s.scores[player]

}
func TestGETPlayers(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{
			"Pepper": 20,
			"Floyd":  10,
		},
	}
	server := &PlayerServer{&store}
	t.Run("return Pepper's score", func(t *testing.T) {
		request := getScoreRequestByPlayerName("Pepper")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		// // PlayerServer(response, request)

		// got := response.Body.String()
		// want := "20"
		assertStringsEqual(t, response.Body.String(), "20")
	})

	t.Run("return Floyd's score", func(t *testing.T) {
		request := getScoreRequestByPlayerName("Floyd")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		// PlayerServer(response, request)

		// got := response.Body.String()
		// want := "10"
		assertStringsEqual(t, response.Body.String(), "10")
	})
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
