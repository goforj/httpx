//go:build ignore
// +build ignore

package main

import (
	"github.com/goforj/httpx"
	"github.com/imroc/req/v3"
)

func main() {
	// WithRetryHook registers a retry hook for the client.

	// Example: hook on retry
	c := httpx.New(httpx.WithRetryHook(func(_ *req.Response, _ error) {}))
	_ = c
}
