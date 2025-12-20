//go:build ignore
// +build ignore

package main

import (
	"bytes"
	"github.com/goforj/httpx"
	"io"
	"net/http"
	"net/http/httptest"
	"time"
)

func main() {
	// UploadProgress enables a default progress spinner and bar for uploads.

	// Example: upload with automatic progress
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		buf := make([]byte, 32*1024)
		for {
			n, err := r.Body.Read(buf)
			if n > 0 {
				time.Sleep(10 * time.Millisecond)
			}
			if err == io.EOF {
				break
			}
			if err != nil {
				return
			}
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()

	payload := bytes.Repeat([]byte("x"), 4*1024*1024)
	pr, pw := io.Pipe()
	go func() {
		defer pw.Close()
		chunk := 64 * 1024
		for i := 0; i < len(payload); i += chunk {
			end := i + chunk
			if end > len(payload) {
				end = len(payload)
			}
			_, _ = pw.Write(payload[i:end])
			time.Sleep(50 * time.Millisecond)
		}
	}()

	c := httpx.New()
	_ = httpx.Post[any, string](c, srv.URL+"/upload", nil,
		httpx.FileReader("file", "payload.bin", pr),
		httpx.UploadProgress(),
	)
}
