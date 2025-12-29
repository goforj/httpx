//go:build ignore
// +build ignore

package main

import (
	"errors"
	"github.com/goforj/httpx"
)

func main() {
	// Error returns a short, human-friendly summary of the HTTP error.

	// Example: check for HTTP errors
	type User struct {
		Name string `json:"name"`
	}

	c := httpx.New()
	res, err := httpx.Get[map[string]any](c, "https://httpbin.org/status/404")
	httpx.Dump(res) // dumps map[string]any
	// map[string]interface {}(nil)
	var httpErr *httpx.HTTPError
	if errors.As(err, &httpErr) {
		_ = httpErr.StatusCode
	}
}
