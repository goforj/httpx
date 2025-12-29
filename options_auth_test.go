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
	_, err := Get[string](c, srv.URL, Auth("Token", "abc123"))
	if err != nil {
		t.Fatalf("auth request failed: %v", err)
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
	_, err := Get[string](c, srv.URL, Bearer("token"))
	if err != nil {
		t.Fatalf("bearer request failed: %v", err)
	}
	_, err = Get[string](c, srv.URL, Basic("user", "pass"))
	if err != nil {
		t.Fatalf("basic request failed: %v", err)
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
