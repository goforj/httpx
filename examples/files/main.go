//go:build ignore
// +build ignore

package main

import "github.com/goforj/httpx"

func main() {
	// Files attaches multiple files from disk as multipart form data.

	// Example: upload multiple files
	c := httpx.New()
	res, _ := httpx.Post[any, map[string]any](c, "https://httpbin.org/post", nil, httpx.Files(map[string]string{
		"fileA": "/tmp/a.txt",
		"fileB": "/tmp/b.txt",
	}))
	httpx.Dump(res) // dumps map[string]any
	// #map[string]interface {} {
	//   files => #map[string]interface {} {
	//     fileA => "hello" #string
	//     fileB => "world" #string
	//   }
	// }
}
