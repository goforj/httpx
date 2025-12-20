package httpx

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"
	"unsafe"

	"github.com/imroc/req/v3"
)

func debugServer() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("ok"))
	}))
}

func TestDumpTo(t *testing.T) {
	buf := &bytes.Buffer{}
	srv := debugServer()
	defer srv.Close()

	c := New()
	res := Get[string](c, srv.URL, DumpTo(buf))
	if res.Err != nil {
		t.Fatalf("request failed: %v", res.Err)
	}
	if buf.Len() == 0 {
		t.Fatalf("expected dump output")
	}
}

func TestDumpToFile(t *testing.T) {
	srv := debugServer()
	defer srv.Close()

	path := t.TempDir() + "/dump.txt"
	c := New()
	res := Get[string](c, srv.URL, DumpToFile(path))
	if res.Err != nil {
		t.Fatalf("request failed: %v", res.Err)
	}
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read dump: %v", err)
	}
	if len(data) == 0 {
		t.Fatalf("expected dump output")
	}
}

func TestDumpAndClientDump(t *testing.T) {
	srv := debugServer()
	defer srv.Close()

	c := New(WithDumpAll(), WithDumpEachRequest())
	res := Get[string](c, srv.URL, Dump())
	if res.Err != nil {
		t.Fatalf("request failed: %v", res.Err)
	}
	if res.Response == nil || res.Response.Dump() == "" {
		t.Fatalf("expected response dump")
	}
}

func TestWithDumpEachRequestTo(t *testing.T) {
	buf := &bytes.Buffer{}
	srv := debugServer()
	defer srv.Close()

	c := New(WithDumpEachRequestTo(buf))
	res := Get[string](c, srv.URL)
	if res.Err != nil {
		t.Fatalf("request failed: %v", res.Err)
	}
	if buf.Len() == 0 {
		t.Fatalf("expected dump output")
	}
}

func TestWithDumpEachRequestToNilAndRespNil(t *testing.T) {
	c := New(WithDumpEachRequestTo(nil))
	buf := &bytes.Buffer{}
	WithDumpEachRequestTo(buf)(c)

	clientVal := reflect.ValueOf(c.Req()).Elem()
	afterField := clientVal.FieldByName("afterResponse")
	afterField = reflect.NewAt(afterField.Type(), unsafe.Pointer(afterField.UnsafeAddr())).Elem()
	if afterField.Len() == 0 {
		t.Fatalf("expected afterResponse middleware")
	}
	mw := afterField.Index(afterField.Len() - 1).Interface().(req.ResponseMiddleware)
	if err := mw(c.Req(), nil); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}
