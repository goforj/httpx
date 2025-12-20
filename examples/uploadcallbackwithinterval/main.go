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
	"time"
)

func main() {
	// UploadCallbackWithInterval registers a callback for upload progress with a minimum interval.

	// Example: throttle upload progress updates
	srv1 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		buf := make([]byte, 32*1024)
		for {
			n, err := r.Body.Read(buf)
			if n > 0 {
				time.Sleep(20 * time.Millisecond)
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
	defer srv1.Close()

	payload1 := bytes.Repeat([]byte("x"), 4*1024*1024)
	pr1, pw1 := io.Pipe()
	go func() {
		defer pw1.Close()
		chunk := 64 * 1024
		for i := 0; i < len(payload1); i += chunk {
			end := i + chunk
			if end > len(payload1) {
				end = len(payload1)
			}
			_, _ = pw1.Write(payload1[i:end])
			time.Sleep(50 * time.Millisecond)
		}
	}()
	c1 := httpx.New()
	total1 := float64(len(payload1))
	_ = httpx.Post[any, string](c1, srv1.URL+"/upload", nil,
		httpx.FileReader("file", "payload.bin", pr1),
		httpx.UploadCallbackWithInterval(func(info req.UploadInfo) {
			percent := float64(info.UploadedSize) / total1 * 100
			fmt.Printf("\rprogress: %.1f%% (%.0f bytes)", percent, float64(info.UploadedSize))
			if info.FileSize > 0 && info.UploadedSize >= info.FileSize {
				fmt.Print("\n")
			}
		}, 200*time.Millisecond),
	)

	// Example: report filename and bytes
	srv2 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		buf := make([]byte, 32*1024)
		for {
			n, err := r.Body.Read(buf)
			if n > 0 {
				time.Sleep(20 * time.Millisecond)
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
	defer srv2.Close()

	payload2 := bytes.Repeat([]byte("x"), 4*1024*1024)
	pr2, pw2 := io.Pipe()
	go func() {
		defer pw2.Close()
		chunk := 64 * 1024
		for i := 0; i < len(payload2); i += chunk {
			end := i + chunk
			if end > len(payload2) {
				end = len(payload2)
			}
			_, _ = pw2.Write(payload2[i:end])
			time.Sleep(50 * time.Millisecond)
		}
	}()
	c2 := httpx.New()
	total2 := float64(len(payload2))
	_ = httpx.Post[any, string](c2, srv2.URL+"/upload", nil,
		httpx.FileReader("file", "payload.bin", pr2),
		httpx.UploadCallbackWithInterval(func(info req.UploadInfo) {
			percent := float64(info.UploadedSize) / total2 * 100
			fmt.Printf("\r%s: %.1f%% (%d bytes)", info.FileName, percent, info.UploadedSize)
			if info.FileSize > 0 && info.UploadedSize >= info.FileSize {
				fmt.Print("\n")
			}
		}, 200*time.Millisecond),
	)
}
