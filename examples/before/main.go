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
	_ = httpx.Get[string](c, "https://example.com", httpx.Before(func(r *req.Request) {
		r.EnableDump()
	}))
}
