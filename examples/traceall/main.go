//go:build ignore
// +build ignore

package main

import "github.com/goforj/httpx"

func main() {
	// TraceAll enables req's client-level trace output for all requests.

	// Example: trace all requests
	c := httpx.New(httpx.TraceAll())
	_ = c
}
