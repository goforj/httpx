package httpx

import (
	"crypto/rand"
	"errors"
	"reflect"
	"strings"
	"testing"

	"github.com/imroc/req/v3"
	"github.com/imroc/req/v3/http2"
)

func TestHTTP2SettingsOption(t *testing.T) {
	c := New(http2SettingsOption())
	if c == nil {
		t.Fatalf("expected client")
	}

	setting := http2.Setting{ID: http2.SettingMaxConcurrentStreams, Val: 100}
	c = New(http2SettingsOption(setting))
	got := getHTTP2Field(t, c, "Settings")
	if got.Len() == 0 {
		t.Fatalf("expected settings to be set")
	}
}

func TestHTTP2ConnectionFlowOption(t *testing.T) {
	c := New(http2ConnectionFlowOption(1))
	got := getHTTP2Field(t, c, "ConnectionFlow")
	if got.Uint() != 1 {
		t.Fatalf("connection flow = %d", got.Uint())
	}
}

func TestHTTP2HeaderPriorityOption(t *testing.T) {
	c := New(http2HeaderPriorityOption(http2.PriorityParam{Weight: 255}))
	got := getHTTP2Field(t, c, "HeaderPriority")
	if got.FieldByName("Weight").Uint() != 255 {
		t.Fatalf("header priority weight = %d", got.FieldByName("Weight").Uint())
	}
}

func TestHTTP2PriorityFramesOption(t *testing.T) {
	c := New(http2PriorityFramesOption())
	got := getHTTP2Field(t, c, "PriorityFrames")
	if got.Len() != 0 {
		t.Fatalf("expected no priority frames, got %d", got.Len())
	}

	c = New(http2PriorityFramesOption(http2.PriorityFrame{StreamID: 3}))
	got = getHTTP2Field(t, c, "PriorityFrames")
	if got.Len() != 1 {
		t.Fatalf("expected priority frames, got %d", got.Len())
	}
}

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

func TestMergeOptionBuilders(t *testing.T) {
	b := mergeOptionBuilders(Header("X-Trace", "1"), Query("q", "go"))
	if len(b.ops) != 2 {
		t.Fatalf("expected 2 options, got %d", len(b.ops))
	}

	c := New(b)
	if got := c.req.Headers.Get("X-Trace"); got != "1" {
		t.Fatalf("client header = %q", got)
	}

	r := req.C().R()
	b.applyRequest(r)
	if got := r.Headers.Get("X-Trace"); got != "1" {
		t.Fatalf("request header = %q", got)
	}
	if got := r.QueryParams.Get("q"); got != "go" {
		t.Fatalf("query param = %q", got)
	}
}

func TestWebkitMultipartBoundaryFunc(t *testing.T) {
	got := webkitMultipartBoundaryFunc()
	prefix := "----WebKitFormBoundary"
	if !strings.HasPrefix(got, prefix) {
		t.Fatalf("boundary = %q", got)
	}
	if len(got) != len(prefix)+16 {
		t.Fatalf("boundary length = %d", len(got))
	}
}

func TestWebkitMultipartBoundaryFuncRandError(t *testing.T) {
	original := rand.Reader
	rand.Reader = failingReader{}
	defer func() {
		rand.Reader = original
	}()
	defer func() {
		if recover() == nil {
			t.Fatalf("expected panic")
		}
	}()
	_ = webkitMultipartBoundaryFunc()
}

func TestFirefoxMultipartBoundaryFunc(t *testing.T) {
	got := firefoxMultipartBoundaryFunc()
	prefix := "-------------------------"
	if !strings.HasPrefix(got, prefix) {
		t.Fatalf("boundary = %q", got)
	}
	if len(got) <= len(prefix) {
		t.Fatalf("boundary length = %d", len(got))
	}
	for i := len(prefix); i < len(got); i++ {
		if got[i] < '0' || got[i] > '9' {
			t.Fatalf("boundary suffix = %q", got[len(prefix):])
		}
	}
}

func TestFirefoxMultipartBoundaryFuncRandError(t *testing.T) {
	original := rand.Reader
	rand.Reader = failingReader{}
	defer func() {
		rand.Reader = original
	}()
	defer func() {
		if recover() == nil {
			t.Fatalf("expected panic")
		}
	}()
	_ = firefoxMultipartBoundaryFunc()
}

type failingReader struct{}

func (failingReader) Read(_ []byte) (int, error) {
	return 0, errors.New("rand error")
}

func getHTTP2Field(t *testing.T, c *Client, name string) reflect.Value {
	t.Helper()
	t2 := reflect.ValueOf(c.req.Transport).Elem().FieldByName("t2")
	if !t2.IsValid() || t2.IsNil() {
		t.Fatalf("expected http2 transport to be present")
	}
	return t2.Elem().FieldByName(name)
}
