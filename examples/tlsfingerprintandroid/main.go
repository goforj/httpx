//go:build ignore
// +build ignore

package main

import "github.com/goforj/httpx"

func main() {
	// TLSFingerprintAndroid applies the Android TLS fingerprint preset.

	// Example: apply Android TLS fingerprint
	c := httpx.New(httpx.TLSFingerprintAndroid())
	_ = c
}
