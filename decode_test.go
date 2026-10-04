package httpx

import (
	"net/http"
	"net/http/httptest"
	"reflect"
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

func TestRawResponsesWithNamedTypes(t *testing.T) {
	type text string
	type bytes []byte
	type octet byte
	type octets []octet

	for _, body := range []string{"hello\x00\xff", ""} {
		name := "nonempty"
		if body == "" {
			name = "empty"
		}
		t.Run(name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				_, _ = w.Write([]byte(body))
			}))
			t.Cleanup(server.Close)

			t.Run("string", func(t *testing.T) {
				testRawResponseType(t, server.URL, text(body))
			})
			t.Run("bytes", func(t *testing.T) {
				testRawResponseType(t, server.URL, bytes(body))
			})
			t.Run("octets", func(t *testing.T) {
				want := make([]octet, len(body))
				for i := range want {
					want[i] = octet(body[i])
				}
				testRawResponseType(t, server.URL, want)
			})
			t.Run("named_octets", func(t *testing.T) {
				want := make(octets, len(body))
				for i := range want {
					want[i] = octet(body[i])
				}
				testRawResponseType(t, server.URL, want)
			})
		})
	}
}

func testRawResponseType[T any](t *testing.T, url string, want T) {
	t.Helper()
	client := New()
	tests := []struct {
		name string
		call func() (T, error)
	}{
		{"function", func() (T, error) { return Get[T](client, url) }},
		{"method", func() (T, error) { return client.Get[T](url) }},
		{"do", func() (T, error) {
			r := client.Raw().R().SetURL(url)
			r.Method = http.MethodGet
			out, _, err := Do[T](r)
			return out, err
		}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := test.call()
			if err != nil {
				t.Fatalf("request failed: %v", err)
			}
			if !reflect.DeepEqual(got, want) {
				t.Fatalf("response = %#v, want %#v", got, want)
			}
		})
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
