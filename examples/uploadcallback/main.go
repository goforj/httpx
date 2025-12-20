//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"github.com/goforj/httpx"
	"github.com/imroc/req/v3"
)

func main() {
	// UploadCallback registers a callback for upload progress.

	// Example: track upload progress
	c := httpx.New()
	_ = httpx.Post[any, string](c, "https://example.com/upload", nil,
		httpx.File("file", "/tmp/report.bin"),
		httpx.UploadCallback(func(info req.UploadInfo) {
			percent := float64(info.UploadedSize) / float64(info.FileSize) * 100
			fmt.Printf("\rprogress: %.1f%%", percent)
			if info.FileSize > 0 && info.UploadedSize >= info.FileSize {
				fmt.Print("\n")
			}
		}),
	)
}
