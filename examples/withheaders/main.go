//go:build ignore
// +build ignore

package main

import "github.com/goforj/httpx"

func main() {
	// WithHeaders sets default headers for all requests.

	// Example: client headers
	c := httpx.New(httpx.WithHeaders(map[string]string{
		"X-Trace": "1",
		"Accept":  "application/json",
	}))
	_ = c
}
