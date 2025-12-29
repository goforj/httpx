//go:build ignore
// +build ignore

package main

import "github.com/goforj/httpx"

func main() {
	// AsFirefox applies the Firefox browser profile (headers including User-Agent, TLS, and HTTP/2 behavior).

	// Example: use a Firefox profile
	c := httpx.New(httpx.AsFirefox())
	res, err := httpx.Get[map[string]any](c, "https://httpbin.org/headers")
	_ = err
	httpx.Dump(res) // dumps map[string]any
	// #map[string]interface {} {
	//   headers => #map[string]interface {} {
	//     User-Agent => "<user-agent>" #string
	//   }
	// }
}
