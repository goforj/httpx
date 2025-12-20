//go:build ignore
// +build ignore

package main

import "github.com/goforj/httpx"

func main() {
	// UploadProgress enables a default progress spinner and bar for uploads.

	// Example: upload with automatic progress
	c := httpx.New()
	_ = httpx.Post[any, string](c, "https://example.com/upload", nil,
		httpx.File("file", "/tmp/report.bin"),
		httpx.UploadProgress(),
	)
}
