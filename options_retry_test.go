package httpx

import (
	"testing"
	"time"

	"github.com/imroc/req/v3"
)

func TestRetryRequestOptions(t *testing.T) {
	r := req.C().R()
	RetryCount(2).applyRequest(r)
	RetryFixedInterval(10 * time.Millisecond).applyRequest(r)
	RetryBackoff(5*time.Millisecond, 50*time.Millisecond).applyRequest(r)
	RetryInterval(func(_ *req.Response, attempt int) time.Duration {
		return time.Duration(attempt) * time.Millisecond
	}).applyRequest(r)
	RetryCondition(func(resp *req.Response, _ error) bool {
		return resp != nil && resp.StatusCode == 503
	}).applyRequest(r)
	RetryHook(func(_ *req.Response, _ error) {}).applyRequest(r)
}

func TestRetryClientOptions(t *testing.T) {
	c := New()
	called := false
	Retry(nil).applyClient(c)
	Retry(func(rc *req.Client) {
		called = true
		rc.SetCommonRetryCount(1)
	}).applyClient(c)
	if !called {
		t.Fatalf("expected retry option to be applied")
	}
	RetryCount(2).applyClient(c)
	RetryFixedInterval(10 * time.Millisecond).applyClient(c)
	RetryBackoff(5*time.Millisecond, 50*time.Millisecond).applyClient(c)
	RetryInterval(func(_ *req.Response, attempt int) time.Duration {
		return time.Duration(attempt) * time.Millisecond
	}).applyClient(c)
	RetryCondition(func(resp *req.Response, _ error) bool {
		return resp != nil && resp.StatusCode == 503
	}).applyClient(c)
	RetryHook(func(_ *req.Response, _ error) {}).applyClient(c)
}
