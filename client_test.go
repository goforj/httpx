package httpx

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/imroc/req/v3"
)

type user struct {
	Name string `json:"name"`
}

type createUser struct {
	Name string `json:"name"`
}

func TestGetTypedSuccess(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/user" {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		_ = json.NewEncoder(w).Encode(user{Name: "roc"})
	}))
	t.Cleanup(server.Close)

	client := New()
	res := Get[user](client, server.URL+"/user")
	if res.Err != nil {
		t.Fatalf("unexpected error: %v", res.Err)
	}
	if res.Body.Name != "roc" {
		t.Fatalf("unexpected response: %s", res.Body.Name)
	}
}

func TestPostTypedSuccess(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var payload createUser
		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		_ = json.NewEncoder(w).Encode(user{Name: payload.Name})
	}))
	t.Cleanup(server.Close)

	client := New()
	res := Post[createUser, user](client, server.URL, createUser{Name: "chris"})
	if res.Err != nil {
		t.Fatalf("unexpected error: %v", res.Err)
	}
	if res.Body.Name != "chris" {
		t.Fatalf("unexpected response: %s", res.Body.Name)
	}
}

func TestRawString(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("plain"))
	}))
	t.Cleanup(server.Close)

	client := New()
	res := Get[string](client, server.URL)
	if res.Err != nil {
		t.Fatalf("unexpected error: %v", res.Err)
	}
	if res.Body != "plain" {
		t.Fatalf("unexpected response: %s", res.Body)
	}
}

func TestHTTPError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("bad"))
	}))
	t.Cleanup(server.Close)

	client := New()
	res := Get[user](client, server.URL)
	if res.Err == nil {
		t.Fatal("expected error")
	}

	var httpErr *HTTPError
	if !errors.As(res.Err, &httpErr) {
		t.Fatalf("expected HTTPError, got %T", res.Err)
	}
	if httpErr.StatusCode != http.StatusBadRequest {
		t.Fatalf("unexpected status: %d", httpErr.StatusCode)
	}
	if string(httpErr.Body) != "bad" {
		t.Fatalf("unexpected body: %s", string(httpErr.Body))
	}
}

func TestErrorMapper(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
	}))
	t.Cleanup(server.Close)

	sentinel := errors.New("mapped")
	client := New(WithErrorMapper(func(resp *req.Response) error {
		return sentinel
	}))

	res := Get[user](client, server.URL)
	if !errors.Is(res.Err, sentinel) {
		t.Fatalf("expected mapped error, got %v", res.Err)
	}
}

func TestOptionsApplied(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/users/123" {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		if r.URL.Query().Get("q") != "ok" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if r.Header.Get("X-Test") != "1" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		_ = json.NewEncoder(w).Encode(user{Name: "ok"})
	}))
	t.Cleanup(server.Close)

	client := New()
	res := Get[user](
		client,
		server.URL+"/users/{id}",
		Path("id", 123),
		Query("q", "ok"),
		Header("X-Test", "1"),
	)
	if res.Err != nil {
		t.Fatalf("unexpected error: %v", res.Err)
	}
	if res.Body.Name != "ok" {
		t.Fatalf("unexpected response: %s", res.Body.Name)
	}
}

func TestResultResponse(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Test", "1")
		_ = json.NewEncoder(w).Encode(user{Name: "ok"})
	}))
	t.Cleanup(server.Close)

	client := New()
	res := Get[user](client, server.URL)
	if res.Err != nil {
		t.Fatalf("unexpected error: %v", res.Err)
	}
	if res.Response == nil || res.Response.Response == nil {
		t.Fatal("expected response")
	}
	if got := res.Response.Header.Get("X-Test"); got != "1" {
		t.Fatalf("unexpected header: %s", got)
	}
	if res.Body.Name != "ok" {
		t.Fatalf("unexpected body: %s", res.Body.Name)
	}
}

func TestEmptySliceOnEmptyBody(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	t.Cleanup(server.Close)

	client := New()
	res := Get[[]user](client, server.URL)
	if res.Err != nil {
		t.Fatalf("unexpected error: %v", res.Err)
	}
	if res.Body == nil {
		t.Fatal("expected empty slice, got nil")
	}
	if len(res.Body) != 0 {
		t.Fatalf("expected empty slice, got %d", len(res.Body))
	}
}
