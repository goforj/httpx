package httpx

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

// genericMethodResponse is the shared response shape for generic method tests.
type genericMethodResponse struct {
	Method string `json:"method"`
}

// TestGenericMethodSurface verifies every verb and context variant uses the receiver client.
func TestGenericMethodSurface(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Method != http.MethodHead {
			_, _ = fmt.Fprintf(w, `{"method":%q}`, r.Method)
		}
	}))
	t.Cleanup(server.Close)

	c := New()
	ctx := context.Background()

	// Generic method expressions and values must remain usable as callbacks.
	get := (*Client).Get[genericMethodResponse]
	getCtx := c.GetCtx[genericMethodResponse]
	_ = (*Client).Post[genericMethodResponse]
	_ = (*Client).Put[genericMethodResponse]
	_ = (*Client).Patch[genericMethodResponse]
	_ = (*Client).Delete[genericMethodResponse]
	_ = (*Client).Head[string]
	_ = (*Client).Options[genericMethodResponse]
	_ = (*Client).PostCtx[genericMethodResponse]
	_ = (*Client).PutCtx[genericMethodResponse]
	_ = (*Client).PatchCtx[genericMethodResponse]
	_ = (*Client).DeleteCtx[genericMethodResponse]
	_ = (*Client).HeadCtx[string]
	_ = (*Client).OptionsCtx[genericMethodResponse]

	tests := []struct {
		name       string
		wantMethod string
		call       func() (genericMethodResponse, error)
	}{
		{"get", http.MethodGet, func() (genericMethodResponse, error) { return get(c, server.URL) }},
		{"post", http.MethodPost, func() (genericMethodResponse, error) { return c.Post[genericMethodResponse](server.URL, nil) }},
		{"put", http.MethodPut, func() (genericMethodResponse, error) {
			return c.Put[genericMethodResponse](server.URL, map[string]string{"name": "Ada"})
		}},
		{"patch", http.MethodPatch, func() (genericMethodResponse, error) {
			return c.Patch[genericMethodResponse](server.URL, map[string]string{"name": "Grace"})
		}},
		{"delete", http.MethodDelete, func() (genericMethodResponse, error) { return c.Delete[genericMethodResponse](server.URL) }},
		{"options", http.MethodOptions, func() (genericMethodResponse, error) { return c.Options[genericMethodResponse](server.URL) }},
		{"get_ctx", http.MethodGet, func() (genericMethodResponse, error) { return getCtx(ctx, server.URL) }},
		{"post_ctx", http.MethodPost, func() (genericMethodResponse, error) { return c.PostCtx[genericMethodResponse](ctx, server.URL, nil) }},
		{"put_ctx", http.MethodPut, func() (genericMethodResponse, error) {
			return c.PutCtx[genericMethodResponse](ctx, server.URL, map[string]string{"name": "Ada"})
		}},
		{"patch_ctx", http.MethodPatch, func() (genericMethodResponse, error) {
			return c.PatchCtx[genericMethodResponse](ctx, server.URL, map[string]string{"name": "Grace"})
		}},
		{"delete_ctx", http.MethodDelete, func() (genericMethodResponse, error) { return c.DeleteCtx[genericMethodResponse](ctx, server.URL) }},
		{"options_ctx", http.MethodOptions, func() (genericMethodResponse, error) { return c.OptionsCtx[genericMethodResponse](ctx, server.URL) }},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := test.call()
			if err != nil {
				t.Fatalf("request failed: %v", err)
			}
			if got.Method != test.wantMethod {
				t.Fatalf("method = %q, want %q", got.Method, test.wantMethod)
			}
		})
	}

	if got, err := c.Head[string](server.URL); err != nil || got != "" {
		t.Fatalf("Head = %q, err %v", got, err)
	}
	if got, err := c.HeadCtx[string](ctx, server.URL); err != nil || got != "" {
		t.Fatalf("HeadCtx = %q, err %v", got, err)
	}
}

// TestGenericContextMethodCancellation verifies context variants forward cancellation to the request.
func TestGenericContextMethodCancellation(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	_, err := New().GetCtx[genericMethodResponse](ctx, "https://example.invalid")
	if err == nil {
		t.Fatal("expected canceled request error")
	}
}
