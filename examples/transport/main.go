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
	_ = c
}
