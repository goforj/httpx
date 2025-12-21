package httpx

import (
	"testing"
	"time"

	"github.com/imroc/req/v3"
)

func TestRetryRequestOptions(t *testing.T) {
	r := req.C().R()
	Opts().RetryCount(2).applyRequest(r)
	Opts().RetryFixedInterval(10 * time.Millisecond).applyRequest(r)
	Opts().RetryBackoff(5*time.Millisecond, 50*time.Millisecond).applyRequest(r)
	Opts().RetryInterval(func(_ *req.Response, attempt int) time.Duration {
		return time.Duration(attempt) * time.Millisecond
	}).applyRequest(r)
	Opts().RetryCondition(func(resp *req.Response, _ error) bool {
		return resp != nil && resp.StatusCode == 503
	}).applyRequest(r)
	Opts().RetryHook(func(_ *req.Response, _ error) {}).applyRequest(r)
}

func TestRetryClientOptions(t *testing.T) {
	c := New()
	called := false
	Opts().Retry(nil).applyClient(c)
	Opts().Retry(func(rc *req.Client) {
		called = true
		rc.SetCommonRetryCount(1)
	}).applyClient(c)
	if !called {
		t.Fatalf("expected retry option to be applied")
	}
	Opts().RetryCount(2).applyClient(c)
	Opts().RetryFixedInterval(10 * time.Millisecond).applyClient(c)
	Opts().RetryBackoff(5*time.Millisecond, 50*time.Millisecond).applyClient(c)
	Opts().RetryInterval(func(_ *req.Response, attempt int) time.Duration {
		return time.Duration(attempt) * time.Millisecond
	}).applyClient(c)
	Opts().RetryCondition(func(resp *req.Response, _ error) bool {
		return resp != nil && resp.StatusCode == 503
	}).applyClient(c)
	Opts().RetryHook(func(_ *req.Response, _ error) {}).applyClient(c)
}
