//go:build ignore
// +build ignore

package main

import (
	"bytes"
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
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		buf := make([]byte, 32*1024)
		for {
			if _, err := r.Body.Read(buf); err != nil {
				if err == io.EOF {
					break
				}
				return
			}
			time.Sleep(10 * time.Millisecond)
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()

	tmp, err := os.CreateTemp("", "httpx-upload-*.bin")
	if err != nil {
		return
	}
	defer os.Remove(tmp.Name())
	payload := bytes.Repeat([]byte("x"), 4*1024*1024)
	_, _ = tmp.Write(payload)
	_ = tmp.Close()

	c := httpx.New()
	_ = httpx.Post[any, string](c, srv.URL+"/upload", nil,
		httpx.File("file", tmp.Name()),
		httpx.UploadProgress(),
	)
}
