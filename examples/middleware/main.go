//go:build ignore
// +build ignore

package main

import (
	"github.com/goforj/httpx"
	"github.com/imroc/req/v3"
)

func main() {
	// Middleware adds request middleware to the client.

	// Example: add request middleware
	c := httpx.New(httpx.Middleware(func(_ *req.Client, r *req.Request) error {
		r.SetHeader("X-Trace", "1")
		return nil
	}))
	res, _ := httpx.Get[map[string]any](c, "https://httpbin.org/headers")
	httpx.Dump(res) // dumps map[string]any
	// #map[string]interface {} {
	//   headers => #map[string]interface {} {
	//     X-Trace => "1" #string
	//   }
	// }
}
