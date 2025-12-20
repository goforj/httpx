//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"github.com/goforj/httpx"
	"github.com/imroc/req/v3"
	"time"
)

func main() {
	// UploadCallbackWithInterval registers a callback for upload progress with a minimum interval.

	// Example: throttle upload progress updates
	c := httpx.New()
	_ = httpx.Post[any, string](c, "https://example.com/upload", nil,
		httpx.File("file", "/tmp/report.bin"),
		httpx.UploadCallbackWithInterval(func(info req.UploadInfo) {
			percent := float64(info.UploadedSize) / float64(info.FileSize) * 100
			fmt.Printf("\rprogress: %.1f%% (%.0f bytes)", percent, float64(info.UploadedSize))
			if info.FileSize > 0 && info.UploadedSize >= info.FileSize {
				fmt.Print("\n")
			}
		}, 200*time.Millisecond),
	)
}
