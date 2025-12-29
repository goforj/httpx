package httpx

import (
	"reflect"
	"strings"
	"testing"

	"github.com/imroc/req/v3"
)

func TestAsChrome(t *testing.T) {
	c := New(AsChrome())
	if c.req.Transport.TLSHandshakeContext == nil {
		t.Fatalf("expected TLS handshake context to be set")
	}
	if got := c.req.Headers.Get("User-Agent"); !strings.Contains(got, "Chrome") {
		t.Fatalf("User-Agent = %q", got)
	}
}

func TestAsFirefox(t *testing.T) {
	c := New(AsFirefox())
	if c.req.Transport.TLSHandshakeContext == nil {
		t.Fatalf("expected TLS handshake context to be set")
	}
	if got := c.req.Headers.Get("User-Agent"); !strings.Contains(got, "Firefox") {
		t.Fatalf("User-Agent = %q", got)
	}
}

func TestAsSafari(t *testing.T) {
	c := New(AsSafari())
	if c.req.Transport.TLSHandshakeContext == nil {
		t.Fatalf("expected TLS handshake context to be set")
	}
	if got := c.req.Headers.Get("User-Agent"); !strings.Contains(got, "Safari") {
		t.Fatalf("User-Agent = %q", got)
	}
}

func TestAsMobile(t *testing.T) {
	c := New(AsMobile())
	if c.req.Transport.TLSHandshakeContext == nil {
		t.Fatalf("expected TLS handshake context to be set")
	}
	if got := c.req.Headers.Get("User-Agent"); !strings.Contains(got, "Android") {
		t.Fatalf("User-Agent = %q", got)
	}
}

func TestAsProfileMethods(t *testing.T) {
	c := New(OptionBuilder{}.AsChrome())
	if got := c.req.Headers.Get("User-Agent"); !strings.Contains(got, "Chrome") {
		t.Fatalf("User-Agent = %q", got)
	}
	c = New(OptionBuilder{}.AsFirefox())
	if got := c.req.Headers.Get("User-Agent"); !strings.Contains(got, "Firefox") {
		t.Fatalf("User-Agent = %q", got)
	}
	c = New(OptionBuilder{}.AsSafari())
	if got := c.req.Headers.Get("User-Agent"); !strings.Contains(got, "Safari") {
		t.Fatalf("User-Agent = %q", got)
	}
	c = New(OptionBuilder{}.AsMobile())
	if got := c.req.Headers.Get("User-Agent"); !strings.Contains(got, "Android") {
		t.Fatalf("User-Agent = %q", got)
	}
}

func TestTLSFingerprintKinds(t *testing.T) {
	kinds := []TLSFingerprintKind{
		TLSFingerprintChromeKind,
		TLSFingerprintFirefoxKind,
		TLSFingerprintSafariKind,
		TLSFingerprintEdgeKind,
		TLSFingerprintAndroidKind,
		TLSFingerprintIOSKind,
		TLSFingerprintRandomizedKind,
	}
	for _, kind := range kinds {
		c := New(TLSFingerprint(kind))
		if c.req.Transport.TLSHandshakeContext == nil {
			t.Fatalf("expected TLS handshake context for kind %v", kind)
		}
	}

	c := New(TLSFingerprint(TLSFingerprintKind(99)))
	if c.req.Transport.TLSHandshakeContext != nil {
		t.Fatalf("expected TLS handshake context to be unset for unknown kind")
	}
}

func TestTLSFingerprintHelpers(t *testing.T) {
	c := New(
		TLSFingerprintChrome(),
		TLSFingerprintFirefox(),
		TLSFingerprintSafari(),
		TLSFingerprintEdge(),
		TLSFingerprintAndroid(),
		TLSFingerprintIOS(),
		TLSFingerprintRandomized(),
	)
	if c.req.Transport.TLSHandshakeContext == nil {
		t.Fatalf("expected TLS handshake context to be set")
	}
}

func TestHeaderOrderPrivate(t *testing.T) {
	empty := headerOrder()
	if len(empty.ops) != 1 {
		t.Fatalf("expected option to be recorded")
	}
	c := New()
	empty.applyClient(c)
	r := req.C().R()
	empty.applyRequest(r)
	if _, ok := r.Headers[req.HeaderOderKey]; ok {
		t.Fatalf("expected no header order to be set")
	}

	b := headerOrder("host", "user-agent")
	c = New(b)
	if c == nil {
		t.Fatalf("expected client")
	}
	r = req.C().R()
	b.applyRequest(r)
	got := r.Headers[req.HeaderOderKey]
	if len(got) != 2 || got[0] != "host" || got[1] != "user-agent" {
		t.Fatalf("header order = %v", got)
	}
}

func TestPseudoHeaderOrderPrivate(t *testing.T) {
	empty := pseudoHeaderOrder()
	if len(empty.ops) != 1 {
		t.Fatalf("expected option to be recorded")
	}
	c := New()
	empty.applyClient(c)
	r := req.C().R()
	empty.applyRequest(r)
	if _, ok := r.Headers[req.PseudoHeaderOderKey]; ok {
		t.Fatalf("expected no pseudo header order to be set")
	}

	b := pseudoHeaderOrder(":method", ":path")
	c = New(b)
	if c == nil {
		t.Fatalf("expected client")
	}
	r = req.C().R()
	b.applyRequest(r)
	got := r.Headers[req.PseudoHeaderOderKey]
	if len(got) != 2 || got[0] != ":method" || got[1] != ":path" {
		t.Fatalf("pseudo header order = %v", got)
	}
}

func TestMultipartBoundaryPrivate(t *testing.T) {
	if got := multipartBoundary(nil); len(got.ops) != 0 {
		t.Fatalf("expected no options, got %d", len(got.ops))
	}
	c := New(multipartBoundary(func() string { return "boundary" }))
	if !hasMultipartBoundaryFunc(c) {
		t.Fatalf("expected multipart boundary func to be set")
	}
}

func hasMultipartBoundaryFunc(c *Client) bool {
	v := reflect.ValueOf(c.req).Elem().FieldByName("multipartBoundaryFunc")
	if !v.IsValid() {
		return false
	}
	return v.Kind() == reflect.Func && !v.IsNil()
}
