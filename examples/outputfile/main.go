//go:build ignore
// +build ignore

package main

import "github.com/goforj/httpx"

func main() {
	// OutputFile streams the response body to a file path.

	// Example: download to file
	c := httpx.New()
	_ = httpx.Get[string](c, "https://example.com/file", httpx.OutputFile("/tmp/file.bin"))
}
