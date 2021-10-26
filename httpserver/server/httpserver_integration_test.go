package server

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRecordWinsAndRetrieveThem(t *testing.T) {
	store := NewInMemoryPlayerStore()
	server := PlayerServer{store}
	player := "Rafael"

	server.ServeHTTP(httptest.NewRecorder(), newPOSTWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPOSTWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPOSTWinRequest(player))

	response := httptest.NewRecorder()
	server.ServeHTTP(response, getScoreRequestByPlayerName(player))
	assertResponseStatusCode(t, response.Code, http.StatusOK)
	assertStringsEqual(t, response.Body.String(), "3")
}
