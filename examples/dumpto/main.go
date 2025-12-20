//go:build ignore
// +build ignore

package main

import (
	"bytes"

	"github.com/goforj/httpx"
)

func main() {
	// DumpTo enables req's request-level dump output to a writer.

	// Example: dump to a buffer
	var buf bytes.Buffer
	c := httpx.New()
	_ = httpx.Get[string](c, "https://example.com", httpx.DumpTo(&buf))
}
