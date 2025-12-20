//go:build ignore
// +build ignore

package main

import (
	"github.com/goforj/httpx"
	"strings"
)

func main() {
	// FileReader attaches a file from a reader as multipart form data.

	// Example: upload from reader
	c := httpx.New()
	_ = httpx.Post[any, string](c, "https://example.com/upload", nil, httpx.FileReader("file", "report.txt", strings.NewReader("hello")))
}
