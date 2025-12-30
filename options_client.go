package httpx

import (
	"net/http"
	"net/url"

	"github.com/imroc/req/v3"
)

// BaseURL sets a base URL on the client.
// @group Client Options
//
// Applies to client configuration only.
// Example: client base URL
//
//	c := httpx.New(httpx.BaseURL("https://httpbin.org"))
//	res, _ := httpx.Get[map[string]any](c, "/uuid")
//	httpx.Dump(res) // dumps map[string]any
//	// #map[string]interface {} {
//	//   uuid => "<uuid>" #string
//	// }
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
//	res, _ := httpx.Get[map[string]any](c, "https://httpbin.org/uuid")
//	httpx.Dump(res) // dumps map[string]any
//	// #map[string]interface {} {
//	//   uuid => "<uuid>" #string
//	// }
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
//	res, _ := httpx.Get[map[string]any](c, "https://httpbin.org/headers")
//	httpx.Dump(res) // dumps map[string]any
//	// #map[string]interface {} {
//	//   headers => #map[string]interface {} {
//	//     X-Trace => "1" #string
//	//   }
//	// }
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
//	res, _ := httpx.Get[map[string]any](c, "https://httpbin.org/status/500")
//	httpx.Dump(res) // dumps map[string]any
//	// map[string]interface {}(nil)
func ErrorMapper(fn ErrorMapperFunc) OptionBuilder {
	return OptionBuilder{}.ErrorMapper(fn)
}

func (b OptionBuilder) ErrorMapper(fn ErrorMapperFunc) OptionBuilder {
	return b.add(clientOnly(func(c *Client) {
		c.errorMapper = fn
	}))
}

// Proxy sets a proxy URL for the client.
// @group Client Options
//
// Applies to client configuration only.
// Example: set proxy URL
//
//	c := httpx.New(httpx.Proxy("http://localhost:8080"))
//	res, _ := httpx.Get[map[string]any](c, "https://httpbin.org/get")
//	httpx.Dump(res) // dumps map[string]any
//	// map[string]interface {}(nil)
func Proxy(proxyURL string) OptionBuilder {
	return OptionBuilder{}.Proxy(proxyURL)
}

func (b OptionBuilder) Proxy(proxyURL string) OptionBuilder {
	return b.add(clientOnly(func(c *Client) {
		if proxyURL == "" {
			return
		}
		c.req.SetProxyURL(proxyURL)
	}))
}

// ProxyFunc sets a proxy function for the client.
// @group Client Options
//
// Applies to client configuration only.
// Example: set proxy function
//
//	c := httpx.New(httpx.ProxyFunc(http.ProxyFromEnvironment))
//	res, _ := httpx.Get[map[string]any](c, "https://httpbin.org/uuid")
//	httpx.Dump(res) // dumps map[string]any
//	// #map[string]interface {} {
//	//   uuid => "<uuid>" #string
//	// }
func ProxyFunc(fn func(*http.Request) (*url.URL, error)) OptionBuilder {
	return OptionBuilder{}.ProxyFunc(fn)
}

func (b OptionBuilder) ProxyFunc(fn func(*http.Request) (*url.URL, error)) OptionBuilder {
	if fn == nil {
		return b
	}
	return b.add(clientOnly(func(c *Client) {
		c.req.SetProxy(fn)
	}))
}

// CookieJar sets the cookie jar for the client.
// @group Client Options
//
// Applies to client configuration only.
// Example: set cookie jar and seed cookies
//
//	jar, _ := cookiejar.New(nil)
//	u, _ := url.Parse("https://httpbin.org")
//	jar.SetCookies(u, []*http.Cookie{
//		{Name: "session", Value: "abc123"},
//	})
//	c := httpx.New(httpx.CookieJar(jar))
//	res, _ := httpx.Get[map[string]any](c, "https://httpbin.org/cookies")
//	httpx.Dump(res) // dumps map[string]any
//	// #map[string]interface {} {
//	//   cookies => #map[string]interface {} {
//	//     session => "abc123" #string
//	//   }
//	// }
func CookieJar(jar http.CookieJar) OptionBuilder {
	return OptionBuilder{}.CookieJar(jar)
}

func (b OptionBuilder) CookieJar(jar http.CookieJar) OptionBuilder {
	return b.add(clientOnly(func(c *Client) {
		c.req.SetCookieJar(jar)
	}))
}

// Redirect sets the redirect policy for the client.
// @group Client Options
//
// Applies to client configuration only.
// Example: disable redirects
//
//	c := httpx.New(httpx.Redirect(req.NoRedirectPolicy()))
//	res, _ := httpx.Get[map[string]any](c, "https://httpbin.org/redirect/1")
//	httpx.Dump(res) // dumps map[string]any
//	// map[string]interface {}(nil)
func Redirect(policies ...req.RedirectPolicy) OptionBuilder {
	return OptionBuilder{}.Redirect(policies...)
}

func (b OptionBuilder) Redirect(policies ...req.RedirectPolicy) OptionBuilder {
	return b.add(clientOnly(func(c *Client) {
		if len(policies) == 0 {
			return
		}
		c.req.SetRedirectPolicy(policies...)
	}))
}
