//go:build ignore
// +build ignore

package main

import "github.com/goforj/httpx"

func main() {
	// TLSFingerprintSafari applies the Safari TLS fingerprint preset.

	// Example: apply Safari TLS fingerprint
	c := httpx.New(httpx.TLSFingerprintSafari())
	_ = c
}
