package httpx

import (
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

func TestOutputFile(t *testing.T) {
	payload := "hello"
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(payload))
	}))
	defer srv.Close()

	c := New()
	path := filepath.Join(t.TempDir(), "out.txt")
	res := Get[string](c, srv.URL, OutputFile(path))
	if res.Err != nil {
		t.Fatalf("download failed: %v", res.Err)
	}
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read output: %v", err)
	}
	if string(data) != payload {
		t.Fatalf("output = %q", string(data))
	}
}
