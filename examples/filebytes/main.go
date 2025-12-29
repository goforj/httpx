//go:build ignore
// +build ignore

package main

import "github.com/goforj/httpx"

func main() {
	// FileBytes attaches a file from bytes as multipart form data.

	// Example: upload bytes as a file
	c := httpx.New()
	res, err := httpx.Post[any, map[string]any](c, "https://httpbin.org/post", nil, httpx.FileBytes("file", "report.txt", []byte("hello")))
	_ = err
	httpx.Dump(res) // dumps map[string]any
	// #map[string]interface {} {
	//   files => #map[string]interface {} {
	//     file => "hello" #string
	//   }
	// }
}
