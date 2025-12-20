package httpx

import (
	"time"

	"github.com/imroc/req/v3"
)

// RetryOption configures retry behavior on the underlying req client.
type RetryOption func(*req.Client)

// RetryCount enables retry for a request and sets the maximum retry count.
// @group Retry
//
// Example: request retry count
//
//	c := httpx.New()
//	_ = httpx.Get[string](c, "https://example.com", httpx.RetryCount(2))
func RetryCount(count int) Option {
	return func(r *req.Request) {
		r.SetRetryCount(count)
	}
}

// RetryFixedInterval sets a fixed retry interval for a request.
// @group Retry
//
// Example: request retry interval
//
//	c := httpx.New()
//	_ = httpx.Get[string](c, "https://example.com", httpx.RetryFixedInterval(200*time.Millisecond))
func RetryFixedInterval(interval time.Duration) Option {
	return func(r *req.Request) {
		r.SetRetryFixedInterval(interval)
	}
}

// RetryBackoff sets a capped exponential backoff retry interval for a request.
// @group Retry
//
// Example: request retry backoff
//
//	c := httpx.New()
//	_ = httpx.Get[string](c, "https://example.com", httpx.RetryBackoff(100*time.Millisecond, 2*time.Second))
func RetryBackoff(min, max time.Duration) Option {
	return func(r *req.Request) {
		r.SetRetryBackoffInterval(min, max)
	}
}

// RetryInterval sets a custom retry interval function for a request.
// @group Retry
//
// Example: custom retry interval
//
//	c := httpx.New()
//	_ = httpx.Get[string](c, "https://example.com", httpx.RetryInterval(func(_ *req.Response, attempt int) time.Duration {
//		return time.Duration(attempt) * 100 * time.Millisecond
//	}))
func RetryInterval(fn req.GetRetryIntervalFunc) Option {
	return func(r *req.Request) {
		r.SetRetryInterval(fn)
	}
}

// RetryCondition sets the retry condition for a request.
// @group Retry
//
// Example: retry on 503
//
//	c := httpx.New()
//	_ = httpx.Get[string](c, "https://example.com", httpx.RetryCondition(func(resp *req.Response, _ error) bool {
//		return resp != nil && resp.StatusCode == 503
//	}))
func RetryCondition(condition req.RetryConditionFunc) Option {
	return func(r *req.Request) {
		r.SetRetryCondition(condition)
	}
}

// RetryHook registers a retry hook for a request.
// @group Retry
//
// Example: hook on retry
//
//	c := httpx.New()
//	_ = httpx.Get[string](c, "https://example.com", httpx.RetryHook(func(_ *req.Response, _ error) {}))
func RetryHook(hook req.RetryHookFunc) Option {
	return func(r *req.Request) {
		r.SetRetryHook(hook)
	}
}

// WithRetry applies a retry configuration to the client.
// @group Retry (Client)
//
// Example: set retry count
//
//	c := httpx.New(httpx.WithRetry(func(rc *req.Client) {
//		rc.SetCommonRetryCount(2)
//	}))
//	_ = c
func WithRetry(opt RetryOption) ClientOption {
	return func(c *Client) {
		if opt != nil {
			opt(c.req)
		}
	}
}

// WithRetryCount enables retry for the client and sets the maximum retry count.
// @group Retry (Client)
//
// Example: client retry count
//
//	c := httpx.New(httpx.WithRetryCount(2))
//	_ = c
func WithRetryCount(count int) ClientOption {
	return func(c *Client) {
		c.req.SetCommonRetryCount(count)
	}
}

// WithRetryFixedInterval sets a fixed retry interval for the client.
// @group Retry (Client)
//
// Example: client retry interval
//
//	c := httpx.New(httpx.WithRetryFixedInterval(200 * time.Millisecond))
//	_ = c
func WithRetryFixedInterval(interval time.Duration) ClientOption {
	return func(c *Client) {
		c.req.SetCommonRetryFixedInterval(interval)
	}
}

// WithRetryBackoff sets a capped exponential backoff retry interval for the client.
// @group Retry (Client)
//
// Example: client retry backoff
//
//	c := httpx.New(httpx.WithRetryBackoff(100*time.Millisecond, 2*time.Second))
//	_ = c
func WithRetryBackoff(min, max time.Duration) ClientOption {
	return func(c *Client) {
		c.req.SetCommonRetryBackoffInterval(min, max)
	}
}

// WithRetryInterval sets a custom retry interval function for the client.
// @group Retry (Client)
//
// Example: client retry interval
//
//	c := httpx.New(httpx.WithRetryInterval(func(_ *req.Response, attempt int) time.Duration {
//		return time.Duration(attempt) * 100 * time.Millisecond
//	}))
//	_ = c
func WithRetryInterval(fn req.GetRetryIntervalFunc) ClientOption {
	return func(c *Client) {
		c.req.SetCommonRetryInterval(fn)
	}
}

// WithRetryCondition sets the retry condition for the client.
// @group Retry (Client)
//
// Example: retry on 503
//
//	c := httpx.New(httpx.WithRetryCondition(func(resp *req.Response, _ error) bool {
//		return resp != nil && resp.StatusCode == 503
//	}))
//	_ = c
func WithRetryCondition(condition req.RetryConditionFunc) ClientOption {
	return func(c *Client) {
		c.req.SetCommonRetryCondition(condition)
	}
}

// WithRetryHook registers a retry hook for the client.
// @group Retry (Client)
//
// Example: hook on retry
//
//	c := httpx.New(httpx.WithRetryHook(func(_ *req.Response, _ error) {}))
//	_ = c
func WithRetryHook(hook req.RetryHookFunc) ClientOption {
	return func(c *Client) {
		c.req.SetCommonRetryHook(hook)
	}
}
