package httpx

import (
	"reflect"
	"strings"
	"testing"

	"github.com/imroc/req/v3"
	"github.com/imroc/req/v3/http2"
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

func TestHTTP2Settings(t *testing.T) {
	c := New(HTTP2Settings())
	if c == nil {
		t.Fatalf("expected client")
	}
	c = New(HTTP2Settings(http2.Setting{ID: http2.SettingMaxConcurrentStreams, Val: 100}))
	if c == nil {
		t.Fatalf("expected client")
	}
}

func TestHTTP2ConnectionFlow(t *testing.T) {
	c := New(HTTP2ConnectionFlow(1))
	if c == nil {
		t.Fatalf("expected client")
	}
}

func TestHTTP2HeaderPriority(t *testing.T) {
	c := New(HTTP2HeaderPriority(http2.PriorityParam{Weight: 255}))
	if c == nil {
		t.Fatalf("expected client")
	}
}

func TestHTTP2PriorityFrames(t *testing.T) {
	c := New(HTTP2PriorityFrames())
	if c == nil {
		t.Fatalf("expected client")
	}
	c = New(HTTP2PriorityFrames(http2.PriorityFrame{StreamID: 3}))
	if c == nil {
		t.Fatalf("expected client")
	}
}

func TestHeaderOrder(t *testing.T) {
	empty := HeaderOrder()
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

	b := HeaderOrder("host", "user-agent")
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

func TestPseudoHeaderOrder(t *testing.T) {
	empty := PseudoHeaderOrder()
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

	b := PseudoHeaderOrder(":method", ":path")
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

func TestMultipartBoundary(t *testing.T) {
	if got := MultipartBoundary(nil); len(got.ops) != 0 {
		t.Fatalf("expected no options, got %d", len(got.ops))
	}
	c := New(MultipartBoundary(func() string { return "boundary" }))
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
