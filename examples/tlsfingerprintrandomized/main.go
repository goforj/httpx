//go:build ignore
// +build ignore

package main

import "github.com/goforj/httpx"

func main() {
	// TLSFingerprintRandomized applies a randomized TLS fingerprint preset.

	// Example: apply randomized TLS fingerprint
	c := httpx.New(httpx.TLSFingerprintRandomized())
	_ = c
}
