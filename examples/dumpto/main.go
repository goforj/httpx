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
	res, _ := httpx.Get[map[string]any](c, "https://httpbin.org/uuid", httpx.DumpTo(&buf))
	httpx.Dump(res) // dumps map[string]any
	// #map[string]interface {} {
	//   uuid => "<uuid>" #string
	// }
}
