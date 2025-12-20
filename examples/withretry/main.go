//go:build ignore
// +build ignore

package main

import (
	"github.com/goforj/httpx"
	"github.com/imroc/req/v3"
)

func main() {
	// WithRetry applies a retry configuration to the client.

	// Example: set retry count
	c := httpx.New(httpx.WithRetry(func(rc *req.Client) {
		rc.SetCommonRetryCount(2)
	}))
	_ = c
}
