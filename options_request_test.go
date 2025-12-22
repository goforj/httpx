package httpx

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"

	"github.com/imroc/req/v3"
)

type requestCapture struct {
	path        string
	headers     http.Header
	query       url.Values
	body        []byte
	contentType string
}

func newCaptureServer(t *testing.T, capture *requestCapture) *httptest.Server {
	t.Helper()
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		capture.path = r.URL.Path
		capture.headers = r.Header.Clone()
		capture.query = r.URL.Query()
		capture.contentType = r.Header.Get("Content-Type")
		body, _ := io.ReadAll(r.Body)
		capture.body = body
		w.WriteHeader(http.StatusOK)
	}))
}

func TestHeaderAndHeaders(t *testing.T) {
	capture := &requestCapture{}
	srv := newCaptureServer(t, capture)
	defer srv.Close()

	c := New()
	res := Get[string](c, srv.URL, Header("X-Trace", "1").Headers(map[string]string{
		"Accept": "application/json",
	}))
	if res.Err != nil {
		t.Fatalf("request failed: %v", res.Err)
	}
	if got := capture.headers.Get("X-Trace"); got != "1" {
		t.Fatalf("X-Trace header = %q", got)
	}
	if got := capture.headers.Get("Accept"); got != "application/json" {
		t.Fatalf("Accept header = %q", got)
	}
}

func TestQueryAndQueries(t *testing.T) {
	capture := &requestCapture{}
	srv := newCaptureServer(t, capture)
	defer srv.Close()

	c := New()
	res := Get[string](c, srv.URL, Query("q", "go").Queries(map[string]string{"ok": "1"}))
	if res.Err != nil {
		t.Fatalf("request failed: %v", res.Err)
	}
	if capture.query.Get("q") != "go" {
		t.Fatalf("q = %q", capture.query.Get("q"))
	}
	if capture.query.Get("ok") != "1" {
		t.Fatalf("ok = %q", capture.query.Get("ok"))
	}
}

func TestQueriesFunctionAlone(t *testing.T) {
	capture := &requestCapture{}
	srv := newCaptureServer(t, capture)
	defer srv.Close()

	c := New()
	res := Get[string](c, srv.URL, Queries(map[string]string{
		"lang": "go",
		"ok":   "1",
	}))
	if res.Err != nil {
		t.Fatalf("request failed: %v", res.Err)
	}
	if capture.query.Get("lang") != "go" {
		t.Fatalf("lang = %q", capture.query.Get("lang"))
	}
	if capture.query.Get("ok") != "1" {
		t.Fatalf("ok = %q", capture.query.Get("ok"))
	}
}

func TestPathAndPaths(t *testing.T) {
	capture := &requestCapture{}
	srv := newCaptureServer(t, capture)
	defer srv.Close()

	c := New()
	res := Get[string](c, srv.URL+"/users/{id}", Path("id", 42))
	if res.Err != nil {
		t.Fatalf("path request failed: %v", res.Err)
	}
	if capture.path != "/users/42" {
		t.Fatalf("path = %q", capture.path)
	}

	capture2 := &requestCapture{}
	srv2 := newCaptureServer(t, capture2)
	defer srv2.Close()

	res = Get[string](c, srv2.URL+"/orgs/{org}/users/{id}", Paths(map[string]any{"org": "goforj", "id": 7}))
	if res.Err != nil {
		t.Fatalf("paths request failed: %v", res.Err)
	}
	if capture2.path != "/orgs/goforj/users/7" {
		t.Fatalf("paths = %q", capture2.path)
	}
}

func TestBodyJSONForm(t *testing.T) {
	capture := &requestCapture{}
	srv := newCaptureServer(t, capture)
	defer srv.Close()

	type payload struct {
		Name string `json:"name"`
	}

	c := New()
	res := Post[any, string](c, srv.URL, nil, Body(payload{Name: "Ana"}))
	if res.Err != nil {
		t.Fatalf("body request failed: %v", res.Err)
	}
	if !strings.Contains(capture.contentType, "application/json") {
		t.Fatalf("content type = %q", capture.contentType)
	}
	var decoded payload
	if err := json.Unmarshal(capture.body, &decoded); err != nil {
		t.Fatalf("decode body: %v", err)
	}
	if decoded.Name != "Ana" {
		t.Fatalf("name = %q", decoded.Name)
	}

	capture2 := &requestCapture{}
	srv2 := newCaptureServer(t, capture2)
	defer srv2.Close()

	res = Post[any, string](c, srv2.URL, nil, JSON(payload{Name: "Bea"}))
	if res.Err != nil {
		t.Fatalf("json request failed: %v", res.Err)
	}
	if !strings.Contains(capture2.contentType, "application/json") {
		t.Fatalf("json content type = %q", capture2.contentType)
	}

	capture3 := &requestCapture{}
	srv3 := newCaptureServer(t, capture3)
	defer srv3.Close()

	res = Post[any, string](c, srv3.URL, nil, Form(map[string]string{"name": "Cam"}))
	if res.Err != nil {
		t.Fatalf("form request failed: %v", res.Err)
	}
	if got := string(capture3.body); !strings.Contains(got, "name=Cam") {
		t.Fatalf("form body = %q", got)
	}
}

func TestTimeoutAndBefore(t *testing.T) {
	capture := &requestCapture{}
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		capture.headers = r.Header.Clone()
		time.Sleep(20 * time.Millisecond)
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()

	c := New()
	res := Get[string](c, srv.URL, Before(func(r *req.Request) {
		r.SetHeader("X-Trace", "1")
	}).Timeout(50*time.Millisecond))
	if res.Err != nil {
		t.Fatalf("request failed: %v", res.Err)
	}
	if got := capture.headers.Get("X-Trace"); got != "1" {
		t.Fatalf("X-Trace header = %q", got)
	}
}

func TestSetBodyVariants(t *testing.T) {
	r := req.C().R()
	setBody(r, nil)

	capture := &requestCapture{}
	srv := newCaptureServer(t, capture)
	defer srv.Close()

	c := New()
	res := Post[any, string](c, srv.URL, nil, Body("hello"))
	if res.Err != nil {
		t.Fatalf("string body request failed: %v", res.Err)
	}
	if got := string(capture.body); got != "hello" {
		t.Fatalf("string body = %q", got)
	}

	capture2 := &requestCapture{}
	srv2 := newCaptureServer(t, capture2)
	defer srv2.Close()
	res = Post[any, string](c, srv2.URL, nil, Body([]byte("bytes")))
	if res.Err != nil {
		t.Fatalf("bytes body request failed: %v", res.Err)
	}
	if got := string(capture2.body); got != "bytes" {
		t.Fatalf("bytes body = %q", got)
	}

	capture3 := &requestCapture{}
	srv3 := newCaptureServer(t, capture3)
	defer srv3.Close()
	res = Post[any, string](c, srv3.URL, nil, Body(bytes.NewBufferString("reader")))
	if res.Err != nil {
		t.Fatalf("reader body request failed: %v", res.Err)
	}
	if got := string(capture3.body); got != "reader" {
		t.Fatalf("reader body = %q", got)
	}
}
