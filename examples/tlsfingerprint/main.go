//go:build ignore
// +build ignore

package main

import "github.com/goforj/httpx/v2"

func main() {
	// TLSFingerprint applies a TLS fingerprint preset.

	// Example: apply a TLS fingerprint preset
	_ = httpx.New(httpx.TLSFingerprint(httpx.TLSFingerprintChromeKind))
}
