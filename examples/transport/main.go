//go:build ignore
// +build ignore

package main

import (
	"github.com/goforj/httpx"
	"net/http"
)

func main() {
	// Transport wraps the underlying transport with a custom RoundTripper.

	// Example: wrap transport
	c := httpx.New(httpx.Transport(http.RoundTripper(http.DefaultTransport)))
	res, _ := httpx.Get[map[string]any](c, "https://httpbin.org/uuid")
	httpx.Dump(res) // dumps map[string]any
	// #map[string]interface {} {
	//   uuid => "<uuid>" #string
	// }
}
