package httpx

import (
	"net/http"
	"time"

	"github.com/imroc/req/v3"
)

// WithBaseURL sets a base URL on the client.
// @group Client Options
//
// Example: client base URL
//
//	c := httpx.New(httpx.WithBaseURL("https://api.example.com"))
//	_ = c
func WithBaseURL(url string) ClientOption {
	return func(c *Client) {
		c.req.SetBaseURL(url)
	}
}

// WithTimeout sets the default timeout for the client.
// @group Client Options
//
// Example: client timeout
//
//	c := httpx.New(httpx.WithTimeout(3 * time.Second))
//	_ = c
func WithTimeout(d time.Duration) ClientOption {
	return func(c *Client) {
		c.req.SetTimeout(d)
	}
}

// WithHeader sets a default header for all requests.
// @group Client Options
//
// Example: client header
//
//	c := httpx.New(httpx.WithHeader("X-Trace", "1"))
//	_ = c
func WithHeader(key, value string) ClientOption {
	return func(c *Client) {
		c.req.SetCommonHeader(key, value)
	}
}

// WithHeaders sets default headers for all requests.
// @group Client Options
//
// Example: client headers
//
//	c := httpx.New(httpx.WithHeaders(map[string]string{
//		"X-Trace": "1",
//		"Accept":  "application/json",
//	}))
//	_ = c
func WithHeaders(values map[string]string) ClientOption {
	return func(c *Client) {
		c.req.SetCommonHeaders(values)
	}
}

// WithTransport wraps the underlying transport with a custom RoundTripper.
// @group Client Options
//
// Example: wrap transport
//
//	c := httpx.New(httpx.WithTransport(http.RoundTripper(http.DefaultTransport)))
//	_ = c
func WithTransport(rt http.RoundTripper) ClientOption {
	return func(c *Client) {
		if rt == nil {
			return
		}
		c.req.Transport.WrapRoundTrip(func(http.RoundTripper) http.RoundTripper {
			return rt
		})
	}
}

// WithMiddleware adds request middleware to the client.
// @group Client Options
//
// Example: add request middleware
//
//	c := httpx.New(httpx.WithMiddleware(func(_ *req.Client, r *req.Request) error {
//		r.SetHeader("X-Trace", "1")
//		return nil
//	}))
//	_ = c
func WithMiddleware(mw ...req.RequestMiddleware) ClientOption {
	return func(c *Client) {
		for _, m := range mw {
			c.req.OnBeforeRequest(m)
		}
	}
}

// WithErrorMapper sets a custom error mapper for non-2xx responses.
// @group Client Options
//
// Example: map error responses
//
//	c := httpx.New(httpx.WithErrorMapper(func(resp *req.Response) error {
//		return fmt.Errorf("status %d", resp.StatusCode)
//	}))
//	_ = c
func WithErrorMapper(fn ErrorMapper) ClientOption {
	return func(c *Client) {
		c.errorMapper = fn
	}
}
