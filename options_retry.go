package httpx

import (
	"time"

	"github.com/imroc/req/v3"
)

// RetryCount enables retry for a request and sets the maximum retry count.
// @group Retry
//
// Example: request retry count
//
//	c := httpx.New()
//	_ = httpx.Get[string](c, "https://example.com", httpx.RetryCount(2))
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
// @group Retry
//
// Example: request retry interval
//
//	c := httpx.New()
//	_ = httpx.Get[string](c, "https://example.com", httpx.RetryFixedInterval(200*time.Millisecond))
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
// @group Retry
//
// Example: request retry backoff
//
//	c := httpx.New()
//	_ = httpx.Get[string](c, "https://example.com", httpx.RetryBackoff(100*time.Millisecond, 2*time.Second))
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
// @group Retry
//
// Example: custom retry interval
//
//	c := httpx.New()
//	_ = httpx.Get[string](c, "https://example.com", httpx.RetryInterval(func(_ *req.Response, attempt int) time.Duration {
//		return time.Duration(attempt) * 100 * time.Millisecond
//	}))
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
// @group Retry
//
// Example: retry on 503
//
//	c := httpx.New()
//	_ = httpx.Get[string](c, "https://example.com", httpx.RetryCondition(func(resp *req.Response, _ error) bool {
//		return resp != nil && resp.StatusCode == 503
//	}))
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
// @group Retry
//
// Example: hook on retry
//
//	c := httpx.New()
//	_ = httpx.Get[string](c, "https://example.com", httpx.RetryHook(func(_ *req.Response, _ error) {}))
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
// @group Retry (Client)
//
// Example: configure client retry
//
//	c := httpx.New(httpx.Retry(func(rc *req.Client) {
//		rc.SetCommonRetryCount(2)
//	}))
//	_ = c
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
