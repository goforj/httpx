//go:build ignore
// +build ignore

package main

import (
	"github.com/goforj/httpx"
	"github.com/imroc/req/v3/http2"
)

func main() {
	// HTTP2HeaderPriority sets the HTTP/2 header priority.

	// Example: customize HTTP/2 header priority
	c := httpx.New(httpx.HTTP2HeaderPriority(http2.PriorityParam{Weight: 255}))
	_ = c
}
