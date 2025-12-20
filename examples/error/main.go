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
	res := httpx.Get[User](c, "https://example.com/users/1")
	var httpErr *httpx.HTTPError
	if errors.As(res.Err, &httpErr) {
		_ = httpErr.StatusCode
	}
}
