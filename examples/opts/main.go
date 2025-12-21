//go:build ignore
// +build ignore

package main

import "github.com/goforj/httpx"

func main() {
	// Opts creates a chainable option builder.

	// Example: chain options
	opt := httpx.Opts().Header("X-Trace", "1").Query("q", "go")
	_ = opt
}
