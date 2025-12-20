//go:build ignore
// +build ignore

package main

import "github.com/goforj/httpx"

func main() {
	// File attaches a file from disk as multipart form data.

	// Example: upload a file
	c := httpx.New()
	_ = httpx.Post[any, string](c, "https://example.com/upload", nil, httpx.File("file", "/tmp/report.txt"))
}
