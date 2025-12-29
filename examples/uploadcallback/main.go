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
	res, err := httpx.Post[any, map[string]any](c, "https://httpbin.org/post", nil,
		httpx.File("file", "/tmp/report.bin"),
		httpx.UploadCallback(func(info req.UploadInfo) {
			percent := float64(info.UploadedSize) / float64(info.FileSize) * 100
			fmt.Printf("\rprogress: %.1f%%", percent)
			if info.FileSize > 0 && info.UploadedSize >= info.FileSize {
				fmt.Print("\n")
			}
		}),
	)
	_ = err
	httpx.Dump(res) // dumps map[string]any
	// #map[string]interface {} {
	//   files => #map[string]interface {} {
	//     file => "<file contents>" #string
	//   }
	// }
}
