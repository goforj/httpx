package httpx

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"
	"time"

	"github.com/imroc/req/v3"
)

func TestWithBaseURLAndHeaders(t *testing.T) {
	var gotPath string
	var gotHeader string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotPath = r.URL.Path
		gotHeader = r.Header.Get("X-Trace")
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()

	c := New(BaseURL(srv.URL).Header("X-Trace", "1"))
	res := Get[string](c, "/users/1")
	if res.Err != nil {
		t.Fatalf("request failed: %v", res.Err)
	}
	if gotPath != "/users/1" {
		t.Fatalf("path = %q", gotPath)
	}
	if gotHeader != "1" {
		t.Fatalf("header = %q", gotHeader)
	}
}

func TestWithHeaders(t *testing.T) {
	var gotAccept string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotAccept = r.Header.Get("Accept")
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()

	c := New(Headers(map[string]string{"Accept": "application/json"}))
	res := Get[string](c, srv.URL)
	if res.Err != nil {
		t.Fatalf("request failed: %v", res.Err)
	}
	if gotAccept != "application/json" {
		t.Fatalf("accept = %q", gotAccept)
	}
}

func TestWithTimeout(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(50 * time.Millisecond)
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()

	c := New(Timeout(10 * time.Millisecond))
	res := Get[string](c, srv.URL)
	if res.Err == nil {
		t.Fatalf("expected timeout error")
	}
}

func TestWithTransport(t *testing.T) {
	called := false
	custom := roundTripperFunc(func(req *http.Request) (*http.Response, error) {
		called = true
		body := io.NopCloser(bytes.NewBufferString("ok"))
		return &http.Response{StatusCode: http.StatusOK, Body: body, Header: make(http.Header)}, nil
	})

	c := New(Transport(custom))
	res := Get[string](c, "https://example.com")
	if res.Err != nil {
		t.Fatalf("request failed: %v", res.Err)
	}
	if !called {
		t.Fatalf("expected custom transport")
	}
}

func TestWithTransportNil(t *testing.T) {
	c := New(Transport(nil))
	if c == nil {
		t.Fatalf("expected client")
	}
}

func TestWithMiddleware(t *testing.T) {
	var got string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		got = r.Header.Get("X-Trace")
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()

	c := New(Middleware(func(_ *req.Client, r *req.Request) error {
		r.SetHeader("X-Trace", "1")
		return nil
	}))
	res := Get[string](c, srv.URL)
	if res.Err != nil {
		t.Fatalf("request failed: %v", res.Err)
	}
	if got != "1" {
		t.Fatalf("header = %q", got)
	}
}

func TestWithErrorMapper(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer srv.Close()

	want := errors.New("mapped error")
	c := New(ErrorMapper(func(_ *req.Response) error {
		return want
	}))
	res := Get[string](c, srv.URL)
	if !errors.Is(res.Err, want) {
		t.Fatalf("expected mapped error")
	}
}

func TestWithProxy(t *testing.T) {
	c := New(Proxy("http://localhost:8080"))
	if c.req.Transport.Options.Proxy == nil {
		t.Fatalf("expected proxy to be set")
	}
}

func TestWithProxyEmpty(t *testing.T) {
	c := New(Proxy(""))
	if c.req.Transport.Options.Proxy == nil {
		t.Fatalf("expected proxy to remain set")
	}
}

func TestWithProxyFunc(t *testing.T) {
	c := New(ProxyFunc(nil))
	if c == nil {
		t.Fatalf("expected client")
	}

	fn := func(req *http.Request) (*url.URL, error) {
		return http.ProxyFromEnvironment(req)
	}
	c = New(ProxyFunc(fn))
	if c.req.Transport.Options.Proxy == nil {
		t.Fatalf("expected proxy func to be set")
	}
}

func TestWithCookieJar(t *testing.T) {
	jar, err := cookiejar.New(nil)
	if err != nil {
		t.Fatalf("cookie jar: %v", err)
	}
	c := New(CookieJar(jar))
	got := reflect.ValueOf(c.req).Elem().FieldByName("httpClient").Elem().FieldByName("Jar")
	if got.IsNil() {
		t.Fatalf("expected cookie jar to be set")
	}
}

func TestWithRedirect(t *testing.T) {
	c := New(Redirect())
	if c == nil {
		t.Fatalf("expected client")
	}

	c = New(Redirect(req.NoRedirectPolicy()))
	got := reflect.ValueOf(c.req).Elem().FieldByName("httpClient").Elem().FieldByName("CheckRedirect")
	if got.IsNil() {
		t.Fatalf("expected redirect policy to be set")
	}
}

type roundTripperFunc func(*http.Request) (*http.Response, error)

func (rt roundTripperFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return rt(req)
}
