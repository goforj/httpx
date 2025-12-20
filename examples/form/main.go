//go:build ignore
// +build ignore

package main

import "github.com/goforj/httpx"

func main() {
	// Form sets form data for the request.

	// Example: submit a form
	c := httpx.New()
	_ = httpx.Post[any, string](c, "https://example.com", nil, httpx.Form(map[string]string{
		"name": "Ana",
	}))
}
