//go:build ignore
// +build ignore

package main

import (
	"bytes"
	"github.com/goforj/httpx"
)

func main() {
	// DumpEachRequestTo enables request-level dumps for each request and writes

	// Example: dump each request to a buffer
	var buf bytes.Buffer
	c := httpx.New(httpx.DumpEachRequestTo(&buf))
	_ = httpx.Get[string](c, "https://example.com")
	_ = buf.String()
}
