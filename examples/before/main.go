//go:build ignore
// +build ignore

package main

import (
	"github.com/goforj/httpx"
	"github.com/imroc/req/v3"
)

func main() {
	// Before runs a hook before the request is sent.

	// Example: mutate req.Request
	c := httpx.New()
	res, _ := httpx.Get[map[string]any](c, "https://httpbin.org/get", httpx.Before(func(r *req.Request) {
		r.EnableDump()
	}))
	httpx.Dump(res) // dumps map[string]any
	// #map[string]interface {} {
	//   url => "https://httpbin.org/get" #string
	// }
}
