package httpx

import "github.com/imroc/req/v3"

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
// @group Advanced
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
// @group Advanced
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
// @group Advanced
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
// @group Advanced
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
// @group Advanced
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
// @group Advanced
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
// @group Advanced
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
// @group Advanced
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

func headerOrder(keys ...string) OptionBuilder {
	return OptionBuilder{}.add(bothOption(
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

func pseudoHeaderOrder(keys ...string) OptionBuilder {
	return OptionBuilder{}.add(bothOption(
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

func multipartBoundary(fn func() string) OptionBuilder {
	if fn == nil {
		return OptionBuilder{}
	}
	return OptionBuilder{}.add(clientOnly(func(c *Client) {
		c.req.SetMultipartBoundaryFunc(fn)
	}))
}
