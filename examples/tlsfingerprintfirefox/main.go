//go:build ignore
// +build ignore

package main

import "github.com/goforj/httpx"

func main() {
	// TLSFingerprintFirefox applies the Firefox TLS fingerprint preset.

	// Example: apply Firefox TLS fingerprint
	c := httpx.New(httpx.TLSFingerprintFirefox())
	_ = c
}
