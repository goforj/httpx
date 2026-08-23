package httpx

import (
	"io"
	"net/http"
	"strings"
	"testing"
)

// benchmarkRoundTripper adapts a function into an in-process HTTP transport.
type benchmarkRoundTripper func(*http.Request) (*http.Response, error)

// benchmarkHTTPResponse is the decoded payload used by the hot-path measurement.
type benchmarkHTTPResponse struct {
	Name string `json:"name"`
}

// RoundTrip executes the benchmark transport callback.
func (roundTrip benchmarkRoundTripper) RoundTrip(request *http.Request) (*http.Response, error) {
	return roundTrip(request)
}

// BenchmarkGenericGet compares the compatibility function and client method without network noise.
func BenchmarkGenericGet(b *testing.B) {
	transport := benchmarkRoundTripper(func(request *http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: http.StatusOK,
			Header:     http.Header{"Content-Type": []string{"application/json"}},
			Body:       io.NopCloser(strings.NewReader(`{"name":"Ada"}`)),
			Request:    request,
		}, nil
	})
	c := New(Transport(transport))

	b.Run("Function", func(b *testing.B) {
		b.ReportAllocs()
		for range b.N {
			if _, err := Get[benchmarkHTTPResponse](c, "https://example.test"); err != nil {
				b.Fatalf("Get: %v", err)
			}
		}
	})
	b.Run("Method", func(b *testing.B) {
		b.ReportAllocs()
		for range b.N {
			if _, err := c.Get[benchmarkHTTPResponse]("https://example.test"); err != nil {
				b.Fatalf("Get: %v", err)
			}
		}
	})
}
