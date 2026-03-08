//go:build ignore
// +build ignore

package main

import "github.com/goforj/httpx/v2"

func main() {
	// TLSFingerprintRandomized applies a randomized TLS fingerprint preset.

	// Example: apply randomized TLS fingerprint
	_ = httpx.New(httpx.TLSFingerprintRandomized())
}
