//go:build ignore
// +build ignore

package main

import "github.com/goforj/httpx"

func main() {
	// File attaches a file from disk as multipart form data.

	// Example: upload a file
	c := httpx.New()
	res, err := httpx.Post[any, map[string]any](c, "https://httpbin.org/post", nil, httpx.File("file", "/tmp/report.txt"))
	_ = err
	httpx.Dump(res) // dumps map[string]any
	// #map[string]interface {} {
	//   files => #map[string]interface {} {
	//     file => "hello" #string
	//   }
	// }
}
