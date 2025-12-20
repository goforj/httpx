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
	"strings"
	"time"
)

func main() {
	// UploadCallback registers a callback for upload progress.

	// Example: track upload progress
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

	srv1 := startServer(10 * time.Millisecond)
	defer srv1.Close()
	path := tempFile(4 * 1024 * 1024)
	if path == "" {
		return
	}
	defer os.Remove(path)

	c := httpx.New()
	barWidth := 20
	spin := []string{"|", "/", "-", "\\"}
	spinIndex := 0
	_ = httpx.Post[any, string](c, srv1.URL+"/upload", nil,
		httpx.File("file", path),
		httpx.UploadCallback(func(info req.UploadInfo) {
			spinIndex = (spinIndex + 1) % len(spin)
			percent := float64(info.UploadedSize) / float64(info.FileSize) * 100
			filled := int(percent / 100 * float64(barWidth))
			bar := strings.Repeat("=", filled) + strings.Repeat("-", barWidth-filled)
			fmt.Printf("\r%s [%s] %.1f%%", spin[spinIndex], bar, percent)
			if info.FileSize > 0 && info.UploadedSize >= info.FileSize {
				fmt.Print("\n")
			}
		}),
	)
}
