//go:build ignore
// +build ignore

package main

import (
	"github.com/goforj/httpx"
	"github.com/imroc/req/v3"
)

func main() {
	// WithRetryCondition sets the retry condition for the client.

	// Example: retry on 503
	c := httpx.New(httpx.WithRetryCondition(func(resp *req.Response, _ error) bool {
		return resp != nil && resp.StatusCode == 503
	}))
	_ = c
}
