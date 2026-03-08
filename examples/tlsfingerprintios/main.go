//go:build ignore
// +build ignore

package main

import "github.com/goforj/httpx/v2"

func main() {
	// TLSFingerprintIOS applies the iOS TLS fingerprint preset.

	// Example: apply iOS TLS fingerprint
	_ = httpx.New(httpx.TLSFingerprintIOS())
}
