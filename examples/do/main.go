//go:build ignore
// +build ignore

package main

import (
	"github.com/goforj/httpx"
	"github.com/imroc/req/v3"
	"net/http"
)

func main() {
	// Do executes a pre-configured req request and returns the decoded body and response.
	// This is the low-level escape hatch when you need full req control.

	// Example: advanced request with response access
	r := req.C().R().SetHeader("X-Trace", "1")
	r.SetURL("https://httpbin.org/headers")
	r.Method = http.MethodGet

	res, rawResp, err := httpx.Do[map[string]any](r)
	if err != nil {
		return
	}
	httpx.Dump(res) // dumps map[string]any
	// #map[string]interface {} {
	//   headers => #map[string]interface {} {
	//     X-Trace => "1" #string
	//   }
	// }
	println(rawResp.StatusCode)
}
