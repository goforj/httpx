package httpx

import (
	"testing"

	"github.com/imroc/req/v3"
)

func TestOptionApplyNilFns(t *testing.T) {
	c := New()
	option{}.applyClient(c)

	r := req.C().R()
	option{}.applyRequest(r)
}

func TestOptionBuilderSkipsNil(t *testing.T) {
	var clientCalled, requestCalled bool
	builder := OptionBuilder{}
	builder = builder.add(option{
		clientFn: func(c *Client) {
			clientCalled = true
		},
		requestFn: func(r *req.Request) {
			requestCalled = true
		},
	})
	builder = builder.add(nil)

	c := New()
	builder.applyClient(c)
	if !clientCalled {
		t.Fatalf("client option not applied")
	}

	r := req.C().R()
	builder.applyRequest(r)
	if !requestCalled {
		t.Fatalf("request option not applied")
	}
}
