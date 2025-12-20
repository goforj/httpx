//go:build ignore
// +build ignore

package main

import "github.com/goforj/httpx"

func main() {
	// FileBytes attaches a file from bytes as multipart form data.

	// Example: upload bytes as a file
	c := httpx.New()
	_ = httpx.Post[any, string](c, "https://example.com/upload", nil, httpx.FileBytes("file", "report.txt", []byte("hello")))
}
