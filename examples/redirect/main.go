//go:build ignore
// +build ignore

package main

import (
	"github.com/goforj/httpx"
	"github.com/imroc/req/v3"
)

func main() {
	// Redirect sets the redirect policy for the client.

	// Example: disable redirects
	c := httpx.New(httpx.Redirect(req.NoRedirectPolicy()))
	_ = c
}
