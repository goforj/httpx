//go:build ignore
// +build ignore

package main

import (
	"github.com/goforj/httpx"
	"github.com/imroc/req/v3"
)

func main() {
	// Retry applies a custom retry configuration to the client.
	// Defaults remain in effect unless the callback modifies them.

	// Example: configure client retry
	_ = httpx.New(httpx.Retry(func(rc *req.Client) {
		rc.SetCommonRetryCount(2)
	}))
}
