//go:build ignore
// +build ignore

package main

import (
	"bytes"
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
			time.Sleep(20 * time.Millisecond)
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
		httpx.UploadCallbackWithInterval(func(info req.UploadInfo) {
			percent := float64(info.UploadedSize) / float64(info.FileSize) * 100
			fmt.Printf("\rprogress: %.1f%% (%.0f bytes)", percent, float64(info.UploadedSize))
			if info.FileSize > 0 && info.UploadedSize >= info.FileSize {
				fmt.Print("\n")
			}
		}, 200*time.Millisecond),
	)
}
