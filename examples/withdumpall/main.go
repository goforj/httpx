//go:build ignore
// +build ignore

package main

import "github.com/goforj/httpx"

func main() {
	// WithDumpAll enables req's client-level dump output for all requests.

	// Example: dump every request and response
	c := httpx.New(httpx.WithDumpAll())
	_ = c
}
