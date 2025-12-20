//go:build ignore
// +build ignore

package main

import "github.com/goforj/httpx"

func main() {
	// DumpToFile enables req's request-level dump output to a file path.

	// Example: dump to a file
	c := httpx.New()
	_ = httpx.Get[string](c, "https://example.com", httpx.DumpToFile("httpx.dump"))
}
