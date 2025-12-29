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
	res, err := httpx.Get[map[string]any](c, "https://httpbin.org/get")
	_ = err
	httpx.Dump(res) // dumps map[string]any
	// #map[string]interface {} {
	//   url => "https://httpbin.org/get" #string
	// }
}
