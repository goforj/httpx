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
	res, err := httpx.Post[any, map[string]any](c, "https://httpbin.org/post", nil, httpx.FileReader("file", "report.txt", strings.NewReader("hello")))
	_ = err
	httpx.Dump(res) // dumps map[string]any
	// #map[string]interface {} {
	//   files => #map[string]interface {} {
	//     file => "hello" #string
	//   }
	// }
}
