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
	res, err := httpx.Post[any, map[string]any](c, "https://httpbin.org/post", nil,
		httpx.File("file", "/tmp/report.bin"),
		httpx.UploadCallbackWithInterval(func(info req.UploadInfo) {
			percent := float64(info.UploadedSize) / float64(info.FileSize) * 100
			fmt.Printf("\rprogress: %.1f%% (%.0f bytes)", percent, float64(info.UploadedSize))
			if info.FileSize > 0 && info.UploadedSize >= info.FileSize {
				fmt.Print("\n")
			}
		}, 200*time.Millisecond),
	)
	_ = err
	httpx.Dump(res) // dumps map[string]any
	// #map[string]interface {} {
	//   files => #map[string]interface {} {
	//     file => "<file contents>" #string
	//   }
	// }
}
