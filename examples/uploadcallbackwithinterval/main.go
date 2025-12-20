//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"github.com/goforj/httpx"
	"github.com/imroc/req/v3"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"time"
)

func main() {
	// UploadCallbackWithInterval registers a callback for upload progress with a minimum interval.

	// Example: throttle upload progress updates
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

	srv := startServer(20 * time.Millisecond)
	defer srv.Close()
	path := tempFile(4 * 1024 * 1024)
	if path == "" {
		return
	}
	defer os.Remove(path)

	c := httpx.New()
	_ = httpx.Post[any, string](c, srv.URL+"/upload", nil,
		httpx.File("file", path),
		httpx.UploadCallbackWithInterval(func(info req.UploadInfo) {
			percent := float64(info.UploadedSize) / float64(info.FileSize) * 100
			fmt.Printf("\rprogress: %.1f%% (%.0f bytes)", percent, float64(info.UploadedSize))
			if info.FileSize > 0 && info.UploadedSize >= info.FileSize {
				fmt.Print("\n")
			}
		}, 200*time.Millisecond),
	)
}
