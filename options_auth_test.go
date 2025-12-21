package httpx

import (
	"encoding/base64"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAuthHeaders(t *testing.T) {
	var gotAuth string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotAuth = r.Header.Get("Authorization")
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()

	c := New()
	res := Get[string](c, srv.URL, Opts().Auth("Token", "abc123"))
	if res.Err != nil {
		t.Fatalf("auth request failed: %v", res.Err)
	}
	if gotAuth != "Token abc123" {
		t.Fatalf("auth header = %q", gotAuth)
	}
}

func TestBearerAndBasic(t *testing.T) {
	var auths []string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auths = append(auths, r.Header.Get("Authorization"))
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()

	c := New()
	res := Get[string](c, srv.URL, Opts().Bearer("token"))
	if res.Err != nil {
		t.Fatalf("bearer request failed: %v", res.Err)
	}
	res = Get[string](c, srv.URL, Opts().Basic("user", "pass"))
	if res.Err != nil {
		t.Fatalf("basic request failed: %v", res.Err)
	}
	if len(auths) != 2 {
		t.Fatalf("auths len = %d", len(auths))
	}
	if auths[0] != "Bearer token" {
		t.Fatalf("bearer header = %q", auths[0])
	}
	wantBasic := "Basic " + base64.StdEncoding.EncodeToString([]byte("user:pass"))
	if auths[1] != wantBasic {
		t.Fatalf("basic header = %q", auths[1])
	}
}
