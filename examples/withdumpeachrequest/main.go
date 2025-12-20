//go:build ignore
// +build ignore

package main

import "github.com/goforj/httpx"

func main() {
	// WithDumpEachRequest enables request-level dumps for each request on the client.

	// Example: dump each request as it is sent
	c := httpx.New(httpx.WithDumpEachRequest())
	_ = c
}
