package httpx

import (
	"net/http"

	"github.com/imroc/req/v3"
)

// BaseURL sets a base URL on the client.
// @group Client Options
//
// Applies to client configuration only.
// Example: client base URL
//
//	c := httpx.New(httpx.BaseURL("https://api.example.com"))
//	_ = c
func BaseURL(url string) OptionBuilder {
	return OptionBuilder{}.BaseURL(url)
}

func (b OptionBuilder) BaseURL(url string) OptionBuilder {
	return b.add(clientOnly(func(c *Client) {
		c.req.SetBaseURL(url)
	}))
}

// Transport wraps the underlying transport with a custom RoundTripper.
// @group Client Options
//
// Applies to client configuration only.
// Example: wrap transport
//
//	c := httpx.New(httpx.Transport(http.RoundTripper(http.DefaultTransport)))
//	_ = c
func Transport(rt http.RoundTripper) OptionBuilder {
	return OptionBuilder{}.Transport(rt)
}

func (b OptionBuilder) Transport(rt http.RoundTripper) OptionBuilder {
	return b.add(clientOnly(func(c *Client) {
		if rt == nil {
			return
		}
		c.req.Transport.WrapRoundTrip(func(http.RoundTripper) http.RoundTripper {
			return rt
		})
	}))
}

// Middleware adds request middleware to the client.
// @group Client Options
//
// Applies only to the client configuration.
// Example: add request middleware
//
//	c := httpx.New(httpx.Middleware(func(_ *req.Client, r *req.Request) error {
//		r.SetHeader("X-Trace", "1")
//		return nil
//	}))
//	_ = c
func Middleware(mw ...req.RequestMiddleware) OptionBuilder {
	return OptionBuilder{}.Middleware(mw...)
}

func (b OptionBuilder) Middleware(mw ...req.RequestMiddleware) OptionBuilder {
	return b.add(clientOnly(func(c *Client) {
		for _, m := range mw {
			c.req.OnBeforeRequest(m)
		}
	}))
}

// ErrorMapper sets a custom error mapper for non-2xx responses.
// @group Client Options
//
// Applies only to the client configuration.
// Example: map error responses
//
//	c := httpx.New(httpx.ErrorMapper(func(resp *req.Response) error {
//		return fmt.Errorf("status %d", resp.StatusCode)
//	}))
//	_ = c
func ErrorMapper(fn ErrorMapperFunc) OptionBuilder {
	return OptionBuilder{}.ErrorMapper(fn)
}

func (b OptionBuilder) ErrorMapper(fn ErrorMapperFunc) OptionBuilder {
	return b.add(clientOnly(func(c *Client) {
		c.errorMapper = fn
	}))
}
