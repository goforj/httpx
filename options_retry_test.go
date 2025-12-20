package httpx

import (
	"testing"
	"time"

	"github.com/imroc/req/v3"
)

func TestRetryRequestOptions(t *testing.T) {
	r := req.C().R()
	RetryCount(2)(r)
	RetryFixedInterval(10 * time.Millisecond)(r)
	RetryBackoff(5*time.Millisecond, 50*time.Millisecond)(r)
	RetryInterval(func(_ *req.Response, attempt int) time.Duration {
		return time.Duration(attempt) * time.Millisecond
	})(r)
	RetryCondition(func(resp *req.Response, _ error) bool {
		return resp != nil && resp.StatusCode == 503
	})(r)
	RetryHook(func(_ *req.Response, _ error) {})(r)
}

func TestRetryClientOptions(t *testing.T) {
	c := New()
	called := false
	WithRetry(nil)(c)
	WithRetry(func(rc *req.Client) {
		called = true
		rc.SetCommonRetryCount(1)
	})(c)
	if !called {
		t.Fatalf("expected retry option to be applied")
	}
	WithRetryCount(2)(c)
	WithRetryFixedInterval(10 * time.Millisecond)(c)
	WithRetryBackoff(5*time.Millisecond, 50*time.Millisecond)(c)
	WithRetryInterval(func(_ *req.Response, attempt int) time.Duration {
		return time.Duration(attempt) * time.Millisecond
	})(c)
	WithRetryCondition(func(resp *req.Response, _ error) bool {
		return resp != nil && resp.StatusCode == 503
	})(c)
	WithRetryHook(func(_ *req.Response, _ error) {})(c)
}
