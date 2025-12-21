//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"github.com/goforj/httpx"
	"github.com/imroc/req/v3"
	"net/http"
	"time"
)

func main() {
	// New creates a client with opinionated defaults and optional overrides.

	// Example: configure all client options
	c := httpx.New(
		httpx.WithBaseURL("https://api.example.com"),
		httpx.WithTimeout(5*time.Second),
		httpx.WithHeader("X-Trace", "1"),
		httpx.WithHeaders(map[string]string{
			"Accept": "application/json",
		}),
		httpx.WithTransport(http.RoundTripper(http.DefaultTransport)),
		httpx.WithMiddleware(func(_ *req.Client, r *req.Request) error {
			r.SetHeader("X-Middleware", "1")
			return nil
		}),
		httpx.WithErrorMapper(func(resp *req.Response) error {
			return fmt.Errorf("status %d", resp.StatusCode)
		}),
	)
	_ = c
}
