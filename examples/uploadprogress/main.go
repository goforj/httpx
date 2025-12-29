//go:build ignore
// +build ignore

package main

import "github.com/goforj/httpx"

func main() {
	// UploadProgress enables a default progress spinner and bar for uploads.

	// Example: upload with automatic progress
	c := httpx.New()
	res, err := httpx.Post[any, map[string]any](c, "https://httpbin.org/post", nil,
		httpx.File("file", "/tmp/report.bin"),
		httpx.UploadProgress(),
	)
	_ = err
	httpx.Dump(res) // dumps map[string]any
	// #map[string]interface {} {
	//   files => #map[string]interface {} {
	//     file => "<file>" #string
	//   }
	// }
}
