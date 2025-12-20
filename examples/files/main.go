//go:build ignore
// +build ignore

package main

import "github.com/goforj/httpx"

func main() {
	// Files attaches multiple files from disk as multipart form data.

	// Example: upload multiple files
	c := httpx.New()
	_ = httpx.Post[any, string](c, "https://example.com/upload", nil, httpx.Files(map[string]string{
		"fileA": "/tmp/a.txt",
		"fileB": "/tmp/b.txt",
	}))
}
