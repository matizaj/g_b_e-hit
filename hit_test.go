package hit

import (
	"net/http"
	"testing"
	"net/http/httptest"
    "sync/atomic"
)

type roundTripFunc func(*http.Request) (*http.Response, error)

func (f roundTripFunc) RoundTrip(r *http.Request) (*http.Response, error) {
	return f(r)
}

func TestSendStatusCode(t *testing.T) {
	t.Parallel()

	req, err := http.NewRequest(http.MethodGet, "/", http.NoBody)
	if err != nil {
		t.Fatalf("creating http request %v", err)
	}

	fake := func(_ *http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: http.StatusInternalServerError,
		}, nil
	}

	client := &http.Client{
		Transport: roundTripFunc(fake),
	}

	result := Send(client, req)
	if result.Status != http.StatusInternalServerError {
		t.Errorf("got %d, want %d", result.Status, http.StatusInternalServerError)
	}

}

func TestSendN(t *testing.T){
	t.Parallel()
	var hits atomic.Int64

	srv := httptest.NewServer(http.HandlerFunc(func(_ http.ResponseWriter, _ *http.Request){
		hits.Add(1)
	}))
	defer srv.Close()

	req, err := http.NewRequest(http.MethodGet, srv.URL, http.NoBody)
	if err != nil {
		t.Fatalf("creating http request %v", err)
	}

	results, err := SendN(t.Context(), 10, req, Options{ Concurrency: 5})
	if err != nil {
		t.Fatalf("SendN() err=%v; want nil", err)
	}

	for range results {
		
	}

	if got := hits.Load(); got !=10 {
		t.Errorf("got %d; want 10", got)
	}
}