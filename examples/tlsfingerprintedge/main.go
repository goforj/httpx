//go:build ignore
// +build ignore

package main

import "github.com/goforj/httpx"

func main() {
	// TLSFingerprintEdge applies the Edge TLS fingerprint preset.

	// Example: apply Edge TLS fingerprint
	c := httpx.New(httpx.TLSFingerprintEdge())
	_ = c
}
