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
	_, err := Get[string](c, srv.URL, DumpTo(buf))
	if err != nil {
		t.Fatalf("request failed: %v", err)
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
	_, err := Get[string](c, srv.URL, DumpToFile(path))
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read dump: %v", err)
	}
	if len(data) == 0 {
		t.Fatalf("expected dump output")
	}
}

func TestDumpAll(t *testing.T) {
	srv := debugServer()
	defer srv.Close()

	c := New(DumpAll())
	_, err := Get[string](c, srv.URL)
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}
}

func TestDumpAndClientDump(t *testing.T) {
	srv := debugServer()
	defer srv.Close()

	req := req.C().R()
	EnableDump().applyRequest(req)
	req.SetURL(srv.URL)
	req.Method = http.MethodGet
	_, resp, err := Do[string](req)
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}
	if resp == nil || resp.Dump() == "" {
		t.Fatalf("expected response dump")
	}
}

func TestWithDumpEachRequestTo(t *testing.T) {
	buf := &bytes.Buffer{}
	srv := debugServer()
	defer srv.Close()

	c := New(DumpEachRequestTo(buf))
	_, err := Get[string](c, srv.URL)
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}
	if buf.Len() == 0 {
		t.Fatalf("expected dump output")
	}
}

func TestWithDumpEachRequestToNilAndRespNil(t *testing.T) {
	c := New(DumpEachRequestTo(nil))
	buf := &bytes.Buffer{}
	DumpEachRequestTo(buf).applyClient(c)

	clientVal := reflect.ValueOf(c.req).Elem()
	afterField := clientVal.FieldByName("afterResponse")
	afterField = reflect.NewAt(afterField.Type(), unsafe.Pointer(afterField.UnsafeAddr())).Elem()
	if afterField.Len() == 0 {
		t.Fatalf("expected afterResponse middleware")
	}
	mw := afterField.Index(afterField.Len() - 1).Interface().(req.ResponseMiddleware)
	if err := mw(c.req, nil); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestDumpEachRequestFunction(t *testing.T) {
	srv := debugServer()
	defer srv.Close()

	c := New(DumpEachRequest())
	_, err := Get[string](c, srv.URL)
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}
}

func TestTrace(t *testing.T) {
	r := req.C().R()
	Trace().applyRequest(r)
	traceField := reflect.ValueOf(r).Elem().FieldByName("trace")
	if traceField.IsNil() {
		t.Fatalf("expected trace to be enabled")
	}
}

func TestTraceAll(t *testing.T) {
	c := New(TraceAll())
	traceField := reflect.ValueOf(c.req).Elem().FieldByName("trace")
	if !traceField.Bool() {
		t.Fatalf("expected trace to be enabled")
	}
}
