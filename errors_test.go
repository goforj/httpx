package httpx

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/imroc/req/v3"
)

func TestHTTPErrorString(t *testing.T) {
	err := (&HTTPError{StatusCode: 400, Status: "Bad Request"}).Error()
	if err == "" {
		t.Fatalf("expected error string")
	}
}

func TestNewHTTPErrorNilResponse(t *testing.T) {
	err := newHTTPError(nil)
	if err.StatusCode != 0 {
		t.Fatalf("expected status 0")
	}
	if err.Status == "" {
		t.Fatalf("expected status message")
	}
}

func TestNewHTTPErrorResponse(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	t.Cleanup(srv.Close)

	resp, reqErr := req.C().R().Get(srv.URL)
	if reqErr != nil {
		t.Fatalf("request failed: %v", reqErr)
	}
	err := newHTTPError(resp)
	if err.StatusCode != 500 {
		t.Fatalf("status code = %d", err.StatusCode)
	}
}
