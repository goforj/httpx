package httpx

import (
	"time"

	"github.com/imroc/req/v3"
)

// RetryCount enables retry for a request and sets the maximum retry count.
// @group Retry
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
func (b OptionBuilder) Retry(fn func(*req.Client)) OptionBuilder {
	return b.add(clientOnly(func(c *Client) {
		if fn == nil {
			return
		}
		fn(c.req)
	}))
}
