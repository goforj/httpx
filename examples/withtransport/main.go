//go:build ignore
// +build ignore

package main

import (
	"net/http"

	"github.com/goforj/httpx"
)

func main() {
	// WithTransport wraps the underlying transport with a custom RoundTripper.

	// Example: wrap transport
	c := httpx.New(httpx.WithTransport(http.RoundTripper(http.DefaultTransport)))
	_ = c
}
