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
	res, err := httpx.Get[map[string]any](c, "https://httpbin.org/redirect/1")
	_ = err
	httpx.Dump(res) // dumps map[string]any
	// #map[string]interface {} {
	//   url => "https://httpbin.org/redirect/1" #string
	// }
}
