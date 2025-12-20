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
	"strings"
	"time"
)

func main() {
	// UploadCallback registers a callback for upload progress.

	// Example: track upload progress
	srv1 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
	barWidth1 := 20
	spin1 := []string{"|", "/", "-", "\\"}
	spinIndex1 := 0
	_ = httpx.Post[any, string](c1, srv1.URL+"/upload", nil,
		httpx.FileReader("file", "payload.bin", pr1),
		httpx.UploadCallback(func(info req.UploadInfo) {
			spinIndex1 = (spinIndex1 + 1) % len(spin1)
			percent := float64(info.UploadedSize) / total1 * 100
			filled := int(percent / 100 * float64(barWidth1))
			bar := strings.Repeat("=", filled) + strings.Repeat("-", barWidth1-filled)
			fmt.Printf("\r%s [%s] %.1f%%", spin1[spinIndex1], bar, percent)
			if info.FileSize > 0 && info.UploadedSize >= info.FileSize {
				fmt.Print("\n")
			}
		}),
	)

	// Example: emit progress percent
	srv2 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
	barWidth2 := 20
	spin2 := []string{"|", "/", "-", "\\"}
	spinIndex2 := 0
	_ = httpx.Post[any, string](c2, srv2.URL+"/upload", nil,
		httpx.FileReader("file", "payload.bin", pr2),
		httpx.UploadCallback(func(info req.UploadInfo) {
			spinIndex2 = (spinIndex2 + 1) % len(spin2)
			percent := float64(info.UploadedSize) / total2 * 100
			filled := int(percent / 100 * float64(barWidth2))
			bar := strings.Repeat("=", filled) + strings.Repeat("-", barWidth2-filled)
			fmt.Printf("\r%s [%s] %.1f%%", spin2[spinIndex2], bar, percent)
			if info.FileSize > 0 && info.UploadedSize >= info.FileSize {
				fmt.Print("\n")
			}
		}),
	)
}
