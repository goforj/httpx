//go:build ignore
// +build ignore

package main

import (
	"github.com/goforj/httpx"
	"github.com/imroc/req/v3/http2"
)

func main() {
	// HTTP2Settings sets HTTP/2 settings frames for the client.

	// Example: customize HTTP/2 settings
	c := httpx.New(httpx.HTTP2Settings(http2.Setting{ID: http2.SettingMaxConcurrentStreams, Val: 100}))
	_ = c
}
