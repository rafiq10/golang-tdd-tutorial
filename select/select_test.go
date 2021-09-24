package selecta

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("check if the faster gets first", func(t *testing.T) {
		slowServer := makeDelayedServer(20 * time.Millisecond)
		fasrServer := makeDelayedServer(0 * time.Millisecond)
		defer slowServer.Close()
		defer fasrServer.Close()

		slowURL := slowServer.URL
		fastURL := fasrServer.URL

		want := fastURL
		got, err := Racer(slowURL, fastURL)

		if err != nil {
			t.Fatalf("did not expect an error but got one: %v", err)
		}

		if got != want {
			t.Errorf("wanted: %v but got: %v", want, got)
		}
	})

	t.Run("returns an error if a server doesn't respond within 10s", func(t *testing.T) {
		server := makeDelayedServer(25 * time.Millisecond)

		defer server.Close()

		_, err := ConfigurableRacer(server.URL, server.URL, 20*time.Millisecond)

		if err == nil {
			t.Error("expected an error but didn't get one")
		}

	})
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		rw.WriteHeader(http.StatusOK)
	}))
}
