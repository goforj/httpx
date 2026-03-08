//go:build ignore
// +build ignore

package main

import "github.com/goforj/httpx/v2"

func main() {
	// TLSFingerprintChrome applies the Chrome TLS fingerprint preset.

	// Example: apply Chrome TLS fingerprint
	_ = httpx.New(httpx.TLSFingerprintChrome())
}
