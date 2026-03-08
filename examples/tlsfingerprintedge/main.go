//go:build ignore
// +build ignore

package main

import "github.com/goforj/httpx/v2"

func main() {
	// TLSFingerprintEdge applies the Edge TLS fingerprint preset.

	// Example: apply Edge TLS fingerprint
	_ = httpx.New(httpx.TLSFingerprintEdge())
}
