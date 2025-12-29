//go:build ignore
// +build ignore

package main

import (
	"github.com/goforj/httpx"
	"github.com/imroc/req/v3/http2"
)

func main() {
	// HTTP2PriorityFrames sets HTTP/2 priority frames for the client.

	// Example: customize HTTP/2 priority frames
	c := httpx.New(httpx.HTTP2PriorityFrames(http2.PriorityFrame{StreamID: 3}))
	_ = c
}
