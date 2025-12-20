//go:build ignore
// +build ignore

package main

import (
	"github.com/goforj/httpx"
	"net/http"
)

func main() {
	// WithTransport wraps the underlying transport with a custom RoundTripper.

	// Example: wrap transport
	c := httpx.New(httpx.WithTransport(http.RoundTripper(http.DefaultTransport)))
	_ = c
}
