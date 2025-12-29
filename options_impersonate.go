package httpx

import (
	"github.com/imroc/req/v3"
	"github.com/imroc/req/v3/http2"
)

// TLSFingerprintKind selects a TLS fingerprint preset.
type TLSFingerprintKind int

const (
	TLSFingerprintChromeKind TLSFingerprintKind = iota
	TLSFingerprintFirefoxKind
	TLSFingerprintSafariKind
	TLSFingerprintEdgeKind
	TLSFingerprintAndroidKind
	TLSFingerprintIOSKind
	TLSFingerprintRandomizedKind
)

// TLSFingerprint applies a TLS fingerprint preset.
// @group TLS Fingerprints
//
// Applies to client configuration only.
// Example: apply a TLS fingerprint preset
//
//	c := httpx.New(httpx.TLSFingerprint(httpx.TLSFingerprintChromeKind))
//	_ = c
func TLSFingerprint(kind TLSFingerprintKind) OptionBuilder {
	return OptionBuilder{}.TLSFingerprint(kind)
}

func (b OptionBuilder) TLSFingerprint(kind TLSFingerprintKind) OptionBuilder {
	return b.add(clientOnly(func(c *Client) {
		switch kind {
		case TLSFingerprintChromeKind:
			c.req.SetTLSFingerprintChrome()
		case TLSFingerprintFirefoxKind:
			c.req.SetTLSFingerprintFirefox()
		case TLSFingerprintSafariKind:
			c.req.SetTLSFingerprintSafari()
		case TLSFingerprintEdgeKind:
			c.req.SetTLSFingerprintEdge()
		case TLSFingerprintAndroidKind:
			c.req.SetTLSFingerprintAndroid()
		case TLSFingerprintIOSKind:
			c.req.SetTLSFingerprintIOS()
		case TLSFingerprintRandomizedKind:
			c.req.SetTLSFingerprintRandomized()
		}
	}))
}

// TLSFingerprintChrome applies the Chrome TLS fingerprint preset.
// @group TLS Fingerprints
//
// Applies to client configuration only.
// Example: apply Chrome TLS fingerprint
//
//	c := httpx.New(httpx.TLSFingerprintChrome())
//	_ = c
func TLSFingerprintChrome() OptionBuilder {
	return OptionBuilder{}.TLSFingerprintChrome()
}

func (b OptionBuilder) TLSFingerprintChrome() OptionBuilder {
	return b.add(clientOnly(func(c *Client) {
		c.req.SetTLSFingerprintChrome()
	}))
}

// TLSFingerprintFirefox applies the Firefox TLS fingerprint preset.
// @group TLS Fingerprints
//
// Applies to client configuration only.
// Example: apply Firefox TLS fingerprint
//
//	c := httpx.New(httpx.TLSFingerprintFirefox())
//	_ = c
func TLSFingerprintFirefox() OptionBuilder {
	return OptionBuilder{}.TLSFingerprintFirefox()
}

func (b OptionBuilder) TLSFingerprintFirefox() OptionBuilder {
	return b.add(clientOnly(func(c *Client) {
		c.req.SetTLSFingerprintFirefox()
	}))
}

// TLSFingerprintSafari applies the Safari TLS fingerprint preset.
// @group TLS Fingerprints
//
// Applies to client configuration only.
// Example: apply Safari TLS fingerprint
//
//	c := httpx.New(httpx.TLSFingerprintSafari())
//	_ = c
func TLSFingerprintSafari() OptionBuilder {
	return OptionBuilder{}.TLSFingerprintSafari()
}

func (b OptionBuilder) TLSFingerprintSafari() OptionBuilder {
	return b.add(clientOnly(func(c *Client) {
		c.req.SetTLSFingerprintSafari()
	}))
}

// TLSFingerprintEdge applies the Edge TLS fingerprint preset.
// @group TLS Fingerprints
//
// Applies to client configuration only.
// Example: apply Edge TLS fingerprint
//
//	c := httpx.New(httpx.TLSFingerprintEdge())
//	_ = c
func TLSFingerprintEdge() OptionBuilder {
	return OptionBuilder{}.TLSFingerprintEdge()
}

func (b OptionBuilder) TLSFingerprintEdge() OptionBuilder {
	return b.add(clientOnly(func(c *Client) {
		c.req.SetTLSFingerprintEdge()
	}))
}

// TLSFingerprintAndroid applies the Android TLS fingerprint preset.
// @group TLS Fingerprints
//
// Applies to client configuration only.
// Example: apply Android TLS fingerprint
//
//	c := httpx.New(httpx.TLSFingerprintAndroid())
//	_ = c
func TLSFingerprintAndroid() OptionBuilder {
	return OptionBuilder{}.TLSFingerprintAndroid()
}

func (b OptionBuilder) TLSFingerprintAndroid() OptionBuilder {
	return b.add(clientOnly(func(c *Client) {
		c.req.SetTLSFingerprintAndroid()
	}))
}

// TLSFingerprintIOS applies the iOS TLS fingerprint preset.
// @group TLS Fingerprints
//
// Applies to client configuration only.
// Example: apply iOS TLS fingerprint
//
//	c := httpx.New(httpx.TLSFingerprintIOS())
//	_ = c
func TLSFingerprintIOS() OptionBuilder {
	return OptionBuilder{}.TLSFingerprintIOS()
}

func (b OptionBuilder) TLSFingerprintIOS() OptionBuilder {
	return b.add(clientOnly(func(c *Client) {
		c.req.SetTLSFingerprintIOS()
	}))
}

// TLSFingerprintRandomized applies a randomized TLS fingerprint preset.
// @group TLS Fingerprints
//
// Applies to client configuration only.
// Example: apply randomized TLS fingerprint
//
//	c := httpx.New(httpx.TLSFingerprintRandomized())
//	_ = c
func TLSFingerprintRandomized() OptionBuilder {
	return OptionBuilder{}.TLSFingerprintRandomized()
}

func (b OptionBuilder) TLSFingerprintRandomized() OptionBuilder {
	return b.add(clientOnly(func(c *Client) {
		c.req.SetTLSFingerprintRandomized()
	}))
}

// HTTP2Settings sets HTTP/2 settings frames for the client.
// @group HTTP2
//
// Applies to client configuration only.
// Example: customize HTTP/2 settings
//
//	c := httpx.New(httpx.HTTP2Settings(http2.Setting{ID: http2.SettingMaxConcurrentStreams, Val: 100}))
//	_ = c
func HTTP2Settings(settings ...http2.Setting) OptionBuilder {
	return OptionBuilder{}.HTTP2Settings(settings...)
}

func (b OptionBuilder) HTTP2Settings(settings ...http2.Setting) OptionBuilder {
	return b.add(clientOnly(func(c *Client) {
		if len(settings) == 0 {
			return
		}
		c.req.SetHTTP2SettingsFrame(settings...)
	}))
}

// HTTP2ConnectionFlow sets the HTTP/2 connection flow control window increment.
// @group HTTP2
//
// Applies to client configuration only.
// Example: customize HTTP/2 connection flow
//
//	c := httpx.New(httpx.HTTP2ConnectionFlow(1_048_576))
//	_ = c
func HTTP2ConnectionFlow(flow uint32) OptionBuilder {
	return OptionBuilder{}.HTTP2ConnectionFlow(flow)
}

func (b OptionBuilder) HTTP2ConnectionFlow(flow uint32) OptionBuilder {
	return b.add(clientOnly(func(c *Client) {
		c.req.SetHTTP2ConnectionFlow(flow)
	}))
}

// HTTP2HeaderPriority sets the HTTP/2 header priority.
// @group HTTP2
//
// Applies to client configuration only.
// Example: customize HTTP/2 header priority
//
//	c := httpx.New(httpx.HTTP2HeaderPriority(http2.PriorityParam{Weight: 255}))
//	_ = c
func HTTP2HeaderPriority(priority http2.PriorityParam) OptionBuilder {
	return OptionBuilder{}.HTTP2HeaderPriority(priority)
}

func (b OptionBuilder) HTTP2HeaderPriority(priority http2.PriorityParam) OptionBuilder {
	return b.add(clientOnly(func(c *Client) {
		c.req.SetHTTP2HeaderPriority(priority)
	}))
}

// HTTP2PriorityFrames sets HTTP/2 priority frames for the client.
// @group HTTP2
//
// Applies to client configuration only.
// Example: customize HTTP/2 priority frames
//
//	c := httpx.New(httpx.HTTP2PriorityFrames(http2.PriorityFrame{StreamID: 3}))
//	_ = c
func HTTP2PriorityFrames(frames ...http2.PriorityFrame) OptionBuilder {
	return OptionBuilder{}.HTTP2PriorityFrames(frames...)
}

func (b OptionBuilder) HTTP2PriorityFrames(frames ...http2.PriorityFrame) OptionBuilder {
	return b.add(clientOnly(func(c *Client) {
		if len(frames) == 0 {
			return
		}
		c.req.SetHTTP2PriorityFrames(frames...)
	}))
}

// HeaderOrder sets the header order for requests.
// @group Impersonation
//
// Applies to client defaults and individual requests.
// Example: set header order
//
//	c := httpx.New(httpx.HeaderOrder("host", "user-agent", "accept"))
//	_ = c
func HeaderOrder(keys ...string) OptionBuilder {
	return OptionBuilder{}.HeaderOrder(keys...)
}

func (b OptionBuilder) HeaderOrder(keys ...string) OptionBuilder {
	return b.add(bothOption(
		func(c *Client) {
			if len(keys) == 0 {
				return
			}
			c.req.SetCommonHeaderOrder(keys...)
		},
		func(r *req.Request) {
			if len(keys) == 0 {
				return
			}
			r.SetHeaderOrder(keys...)
		},
	))
}

// PseudoHeaderOrder sets the HTTP/2 pseudo header order for requests.
// @group Impersonation
//
// Applies to client defaults and individual requests.
// Example: set pseudo header order
//
//	c := httpx.New(httpx.PseudoHeaderOrder(":method", ":authority", ":scheme", ":path"))
//	_ = c
func PseudoHeaderOrder(keys ...string) OptionBuilder {
	return OptionBuilder{}.PseudoHeaderOrder(keys...)
}

func (b OptionBuilder) PseudoHeaderOrder(keys ...string) OptionBuilder {
	return b.add(bothOption(
		func(c *Client) {
			if len(keys) == 0 {
				return
			}
			c.req.SetCommonPseudoHeaderOder(keys...)
		},
		func(r *req.Request) {
			if len(keys) == 0 {
				return
			}
			r.SetPseudoHeaderOrder(keys...)
		},
	))
}

// MultipartBoundary overrides the default multipart boundary generator.
// @group Impersonation
//
// Applies to client configuration only.
// Example: customize multipart boundary
//
//	c := httpx.New(httpx.MultipartBoundary(func() string { return "boundary" }))
//	_ = c
func MultipartBoundary(fn func() string) OptionBuilder {
	return OptionBuilder{}.MultipartBoundary(fn)
}

func (b OptionBuilder) MultipartBoundary(fn func() string) OptionBuilder {
	if fn == nil {
		return b
	}
	return b.add(clientOnly(func(c *Client) {
		c.req.SetMultipartBoundaryFunc(fn)
	}))
}
