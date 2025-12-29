//go:build ignore
// +build ignore

package main

import "github.com/goforj/httpx"

func main() {
	// OutputFile streams the response body to a file path.

	// Example: download to file
	c := httpx.New()
	res, err := httpx.Get[map[string]any](c, "https://httpbin.org/bytes/1024", httpx.OutputFile("/tmp/file.bin"))
	_ = err
	httpx.Dump(res) // dumps map[string]any
	// map[string]interface {}(nil)
}
