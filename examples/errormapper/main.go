//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"github.com/goforj/httpx"
	"github.com/imroc/req/v3"
)

func main() {
	// ErrorMapper sets a custom error mapper for non-2xx responses.

	// Example: map error responses
	c := httpx.New(httpx.ErrorMapper(func(resp *req.Response) error {
		return fmt.Errorf("status %d", resp.StatusCode)
	}))
	res, err := httpx.Get[map[string]any](c, "https://httpbin.org/status/500")
	_ = err
	httpx.Dump(res) // dumps map[string]any
	// map[string]interface {}(nil)
}
