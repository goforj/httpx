//go:build ignore
// +build ignore

package main

import "github.com/goforj/httpx/v2"

func main() {
	// TLSFingerprintFirefox applies the Firefox TLS fingerprint preset.

	// Example: apply Firefox TLS fingerprint
	_ = httpx.New(httpx.TLSFingerprintFirefox())
}
