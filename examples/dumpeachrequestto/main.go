//go:build ignore
// +build ignore

package main

import (
	"bytes"
	"github.com/goforj/httpx"
)

func main() {
	// DumpEachRequestTo enables request-level dumps for each request and writes them to the provided output.

	// Example: dump each request to a buffer
	var buf bytes.Buffer
	c := httpx.New(httpx.DumpEachRequestTo(&buf))
	res, _ := httpx.Get[map[string]any](c, "https://httpbin.org/uuid")
	httpx.Dump(res) // dumps map[string]any
	// #map[string]interface {} {
	//   uuid => "<uuid>" #string
	// }
	_ = buf.String()
}
