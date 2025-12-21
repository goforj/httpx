package httpx

import (
	"net/http"

	"github.com/imroc/req/v3"
)

// BaseURL sets a base URL on the client.
// @group Client Options
func (b OptionBuilder) BaseURL(url string) OptionBuilder {
	return b.add(clientOnly(func(c *Client) {
		c.req.SetBaseURL(url)
	}))
}

// Transport wraps the underlying transport with a custom RoundTripper.
// @group Client Options
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
func (b OptionBuilder) Middleware(mw ...req.RequestMiddleware) OptionBuilder {
	return b.add(clientOnly(func(c *Client) {
		for _, m := range mw {
			c.req.OnBeforeRequest(m)
		}
	}))
}

// ErrorMapper sets a custom error mapper for non-2xx responses.
// @group Client Options
func (b OptionBuilder) ErrorMapper(fn ErrorMapperFunc) OptionBuilder {
	return b.add(clientOnly(func(c *Client) {
		c.errorMapper = fn
	}))
}
