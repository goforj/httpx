//go:build ignore
// +build ignore

package main

import "github.com/goforj/httpx"

func main() {
	// TLSFingerprintChrome applies the Chrome TLS fingerprint preset.

	// Example: apply Chrome TLS fingerprint
	c := httpx.New(httpx.TLSFingerprintChrome())
	_ = c
}
