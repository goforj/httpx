package httpx

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/imroc/req/v3"
)

func TestRawKindAndDecode(t *testing.T) {
	if rawKindOf[string]() != rawString {
		t.Fatalf("expected rawString")
	}
	if rawKindOf[[]byte]() != rawBytes {
		t.Fatalf("expected rawBytes")
	}
	if rawKindOf[[]int]() != rawNone {
		t.Fatalf("expected rawNone for slice")
	}
	if rawKindOf[int]() != rawNone {
		t.Fatalf("expected rawNone")
	}

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("hello"))
	}))
	t.Cleanup(srv.Close)

	resp, err := req.C().R().Get(srv.URL)
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}
	if got := decodeRaw[string](resp); got != "hello" {
		t.Fatalf("decode string = %q", got)
	}
	if got := decodeRaw[[]byte](resp); string(got) != "hello" {
		t.Fatalf("decode bytes = %q", string(got))
	}
	if got := decodeRaw[int](resp); got != 0 {
		t.Fatalf("decode int = %d", got)
	}
}

func TestEnsureNonNil(t *testing.T) {
	var p *int
	ensureNonNil(p)

	var s []int
	ensureNonNil(&s)
	if s == nil {
		t.Fatalf("expected slice initialized")
	}

	var m map[string]int
	ensureNonNil(&m)
	if m == nil {
		t.Fatalf("expected map initialized")
	}
}

func TestIsEmptyBody(t *testing.T) {
	if isEmptyBody(nil) {
		t.Fatalf("expected false for nil response")
	}

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	t.Cleanup(srv.Close)

	resp, err := req.C().R().Get(srv.URL)
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}
	if !isEmptyBody(resp) {
		t.Fatalf("expected empty body")
	}

	srv2 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("data"))
	}))
	t.Cleanup(srv2.Close)

	resp2, err := req.C().R().Get(srv2.URL)
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}
	if isEmptyBody(resp2) {
		t.Fatalf("expected non-empty body")
	}

	srv3 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("   "))
	}))
	t.Cleanup(srv3.Close)

	resp3, err := req.C().R().Get(srv3.URL)
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}
	if !isEmptyBody(resp3) {
		t.Fatalf("expected empty body for whitespace")
	}
}
