package httpx

import (
	"time"

	"github.com/imroc/req/v3"
)

// RetryCount enables retry for a request and sets the maximum retry count.
// Default behavior from req: retries are disabled (count = 0). When enabled,
// retries happen only on request errors unless RetryCondition is set, and the
// default interval is a fixed 100ms between attempts.
// @group Retry
//
// Applies to both client defaults and individual requests.
// Example: retry count
//
//	// Apply to all requests
//	c := httpx.New(httpx.RetryCount(2))
//	res, err := httpx.Get[map[string]any](c, "https://httpbin.org/uuid")
//	_ = err
//	httpx.Dump(res) // dumps map[string]any
//	// #map[string]interface {} {
//	//   uuid => "<uuid>" #string
//	// }
//
//	// Apply to a single request
//	res, err = httpx.Get[map[string]any](c, "https://httpbin.org/uuid", httpx.RetryCount(2))
//	_ = err
//	httpx.Dump(res) // dumps map[string]any
//	// #map[string]interface {} {
//	//   uuid => "<uuid>" #string
//	// }
func RetryCount(count int) OptionBuilder {
	return OptionBuilder{}.RetryCount(count)
}

func (b OptionBuilder) RetryCount(count int) OptionBuilder {
	return b.add(bothOption(
		func(c *Client) {
			c.req.SetCommonRetryCount(count)
		},
		func(r *req.Request) {
			r.SetRetryCount(count)
		},
	))
}

// RetryFixedInterval sets a fixed retry interval for a request.
// Overrides the req default interval (fixed 100ms).
// @group Retry
//
// Applies to both client defaults and individual requests.
// Example: retry interval
//
//	// Apply to all requests
//	c := httpx.New(httpx.RetryFixedInterval(200 * time.Millisecond))
//	res, err := httpx.Get[map[string]any](c, "https://httpbin.org/uuid")
//	_ = err
//	httpx.Dump(res) // dumps map[string]any
//	// #map[string]interface {} {
//	//   uuid => "<uuid>" #string
//	// }
//
//	// Apply to a single request
//	res, err = httpx.Get[map[string]any](c, "https://httpbin.org/uuid", httpx.RetryFixedInterval(200*time.Millisecond))
//	_ = err
//	httpx.Dump(res) // dumps map[string]any
//	// #map[string]interface {} {
//	//   uuid => "<uuid>" #string
//	// }
func RetryFixedInterval(interval time.Duration) OptionBuilder {
	return OptionBuilder{}.RetryFixedInterval(interval)
}

func (b OptionBuilder) RetryFixedInterval(interval time.Duration) OptionBuilder {
	return b.add(bothOption(
		func(c *Client) {
			c.req.SetCommonRetryFixedInterval(interval)
		},
		func(r *req.Request) {
			r.SetRetryFixedInterval(interval)
		},
	))
}

// RetryBackoff sets a capped exponential backoff retry interval for a request.
// Overrides the req default interval (fixed 100ms) with jittered backoff.
// @group Retry
//
// Applies to both client defaults and individual requests.
// Example: retry backoff
//
//	// Apply to all requests
//	c := httpx.New(httpx.RetryBackoff(100*time.Millisecond, 2*time.Second))
//	res, err := httpx.Get[map[string]any](c, "https://httpbin.org/uuid")
//	_ = err
//	httpx.Dump(res) // dumps map[string]any
//	// #map[string]interface {} {
//	//   uuid => "<uuid>" #string
//	// }
//
//	// Apply to a single request
//	res, err = httpx.Get[map[string]any](c, "https://httpbin.org/uuid", httpx.RetryBackoff(100*time.Millisecond, 2*time.Second))
//	_ = err
//	httpx.Dump(res) // dumps map[string]any
//	// #map[string]interface {} {
//	//   uuid => "<uuid>" #string
//	// }
func RetryBackoff(min, max time.Duration) OptionBuilder {
	return OptionBuilder{}.RetryBackoff(min, max)
}

func (b OptionBuilder) RetryBackoff(min, max time.Duration) OptionBuilder {
	return b.add(bothOption(
		func(c *Client) {
			c.req.SetCommonRetryBackoffInterval(min, max)
		},
		func(r *req.Request) {
			r.SetRetryBackoffInterval(min, max)
		},
	))
}

// RetryInterval sets a custom retry interval function for a request.
// Overrides the req default interval (fixed 100ms).
// @group Retry
//
// Applies to both client defaults and individual requests.
// Example: custom retry interval
//
//	// Apply to all requests
//	c := httpx.New(httpx.RetryInterval(func(_ *req.Response, attempt int) time.Duration {
//		return time.Duration(attempt) * 100 * time.Millisecond
//	}))
//	res, err := httpx.Get[map[string]any](c, "https://httpbin.org/uuid")
//	_ = err
//	httpx.Dump(res) // dumps map[string]any
//	// #map[string]interface {} {
//	//   uuid => "<uuid>" #string
//	// }
//
//	// Apply to a single request
//	res, err = httpx.Get[map[string]any](c, "https://httpbin.org/uuid", httpx.RetryInterval(func(_ *req.Response, attempt int) time.Duration {
//		return time.Duration(attempt) * 100 * time.Millisecond
//	}))
//	_ = err
//	httpx.Dump(res) // dumps map[string]any
//	// #map[string]interface {} {
//	//   uuid => "<uuid>" #string
//	// }
func RetryInterval(fn req.GetRetryIntervalFunc) OptionBuilder {
	return OptionBuilder{}.RetryInterval(fn)
}

func (b OptionBuilder) RetryInterval(fn req.GetRetryIntervalFunc) OptionBuilder {
	return b.add(bothOption(
		func(c *Client) {
			c.req.SetCommonRetryInterval(fn)
		},
		func(r *req.Request) {
			r.SetRetryInterval(fn)
		},
	))
}

// RetryCondition sets the retry condition for a request.
// Overrides the default behavior (retry only when a request error occurs).
// @group Retry
//
// Applies to both client defaults and individual requests.
// Example: retry on 503
//
//	// Apply to all requests
//	c := httpx.New(httpx.RetryCondition(func(resp *req.Response, _ error) bool {
//		return resp != nil && resp.StatusCode == 503
//	}))
//	res, err := httpx.Get[map[string]any](c, "https://httpbin.org/status/503")
//	_ = err
//	httpx.Dump(res) // dumps map[string]any
//	// map[string]interface {}(nil)
//
//	// Apply to a single request
//	res, err = httpx.Get[map[string]any](c, "https://httpbin.org/status/503", httpx.RetryCondition(func(resp *req.Response, _ error) bool {
//		return resp != nil && resp.StatusCode == 503
//	}))
//	_ = err
//	httpx.Dump(res) // dumps map[string]any
//	// map[string]interface {}(nil)
func RetryCondition(condition req.RetryConditionFunc) OptionBuilder {
	return OptionBuilder{}.RetryCondition(condition)
}

func (b OptionBuilder) RetryCondition(condition req.RetryConditionFunc) OptionBuilder {
	return b.add(bothOption(
		func(c *Client) {
			c.req.SetCommonRetryCondition(condition)
		},
		func(r *req.Request) {
			r.SetRetryCondition(condition)
		},
	))
}

// RetryHook registers a retry hook for a request.
// Runs before each retry attempt; no hooks are configured by default.
// @group Retry
//
// Applies to both client defaults and individual requests.
// Example: hook on retry
//
//	// Apply to all requests
//	c := httpx.New(httpx.RetryHook(func(_ *req.Response, _ error) {}))
//	res, err := httpx.Get[map[string]any](c, "https://httpbin.org/uuid")
//	_ = err
//	httpx.Dump(res) // dumps map[string]any
//	// #map[string]interface {} {
//	//   uuid => "<uuid>" #string
//	// }
//
//	// Apply to a single request
//	res, err = httpx.Get[map[string]any](c, "https://httpbin.org/uuid", httpx.RetryHook(func(_ *req.Response, _ error) {}))
//	_ = err
//	httpx.Dump(res) // dumps map[string]any
//	// #map[string]interface {} {
//	//   uuid => "<uuid>" #string
//	// }
func RetryHook(hook req.RetryHookFunc) OptionBuilder {
	return OptionBuilder{}.RetryHook(hook)
}

func (b OptionBuilder) RetryHook(hook req.RetryHookFunc) OptionBuilder {
	return b.add(bothOption(
		func(c *Client) {
			c.req.SetCommonRetryHook(hook)
		},
		func(r *req.Request) {
			r.SetRetryHook(hook)
		},
	))
}

// Retry applies a custom retry configuration to the client.
// Defaults remain in effect unless the callback modifies them.
// @group Retry (Client)
//
// Applies only to client configuration.
// Example: configure client retry
//
//	_ = httpx.New(httpx.Retry(func(rc *req.Client) {
//		rc.SetCommonRetryCount(2)
//	}))
func Retry(fn func(*req.Client)) OptionBuilder {
	return OptionBuilder{}.Retry(fn)
}

func (b OptionBuilder) Retry(fn func(*req.Client)) OptionBuilder {
	return b.add(clientOnly(func(c *Client) {
		if fn == nil {
			return
		}
		fn(c.req)
	}))
}
