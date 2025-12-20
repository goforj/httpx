//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"github.com/goforj/httpx"
	"github.com/imroc/req/v3"
)

func main() {
	// WithErrorMapper sets a custom error mapper for non-2xx responses.

	// Example: map error responses
	c := httpx.New(httpx.WithErrorMapper(func(resp *req.Response) error {
		return fmt.Errorf("status %d", resp.StatusCode)
	}))
	_ = c
}
