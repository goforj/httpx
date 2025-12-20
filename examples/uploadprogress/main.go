//go:build ignore
// +build ignore

package main

import (
	"github.com/goforj/httpx"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"time"
)

func main() {
	// UploadProgress enables a default progress spinner and bar for uploads.

	// Example: upload with automatic progress
	startServer := func(delay time.Duration) *httptest.Server {
		return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer r.Body.Close()
			buf := make([]byte, 32*1024)
			for {
				if _, err := r.Body.Read(buf); err != nil {
					if err == io.EOF {
						break
					}
					return
				}
				time.Sleep(delay)
			}
			w.WriteHeader(http.StatusOK)
		}))
	}
	tempFile := func(size int) string {
		tmp, err := os.CreateTemp("", "httpx-upload-*.bin")
		if err != nil {
			return ""
		}
		_, _ = tmp.Write(make([]byte, size))
		_ = tmp.Close()
		return tmp.Name()
	}

	srv := startServer(10 * time.Millisecond)
	defer srv.Close()
	path := tempFile(4 * 1024 * 1024)
	if path == "" {
		return
	}
	defer os.Remove(path)

	c := httpx.New()
	_ = httpx.Post[any, string](c, srv.URL+"/upload", nil,
		httpx.File("file", path),
		httpx.UploadProgress(),
	)
}
