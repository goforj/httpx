//go:build ignore
// +build ignore

package main

import (
	"bytes"
	"fmt"
	"github.com/goforj/httpx"
	"github.com/imroc/req/v3"
	"net/http"
	"time"
)

func main() {
	// New creates a client with opinionated defaults and optional overrides.

	// Example: configure all client options
	var buf bytes.Buffer
	c := httpx.New(httpx.
		BaseURL("https://api.example.com").
		Timeout(5*time.Second).
		Header("X-Trace", "1").
		Headers(map[string]string{
			"Accept": "application/json",
		}).
		Transport(http.RoundTripper(http.DefaultTransport)).
		Middleware(func(_ *req.Client, r *req.Request) error {
			r.SetHeader("X-Middleware", "1")
			return nil
		}).
		ErrorMapper(func(resp *req.Response) error {
			return fmt.Errorf("status %d", resp.StatusCode)
		}).
		DumpAll().
		DumpEachRequest().
		DumpEachRequestTo(&buf).
		Retry(func(rc *req.Client) {
			rc.SetCommonRetryCount(2)
		}).
		RetryCount(2).
		RetryFixedInterval(200*time.Millisecond).
		RetryBackoff(100*time.Millisecond, 2*time.Second).
		RetryInterval(func(_ *req.Response, attempt int) time.Duration {
			return time.Duration(attempt) * 100 * time.Millisecond
		}).
		RetryCondition(func(resp *req.Response, _ error) bool {
			return resp != nil && resp.StatusCode == 503
		}).
		RetryHook(func(_ *req.Response, _ error) {}),
	)
	_ = c
}
