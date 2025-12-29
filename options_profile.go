package httpx

import (
	"crypto/rand"
	"encoding/binary"
	"io"
	"math/big"
	"strconv"
	"strings"

	"github.com/imroc/req/v3/http2"
)

// AsChrome applies the Chrome browser profile (headers, TLS, and HTTP/2 behavior).
// @group Browser Profiles
//
// Applies to client configuration only.
// Example: use a Chrome profile
//
//	c := httpx.New(httpx.AsChrome())
//	_ = c
func AsChrome() OptionBuilder {
	return mergeOptionBuilders(
		TLSFingerprint(TLSFingerprintChromeKind),
		http2SettingsOption(chromeHTTP2Settings...),
		http2ConnectionFlowOption(chromeHTTP2ConnectionFlow),
		pseudoHeaderOrder(chromePseudoHeaderOrder...),
		headerOrder(chromeHeaderOrder...),
		Headers(chromeHeaders),
		http2HeaderPriorityOption(chromeHeaderPriority),
		multipartBoundary(webkitMultipartBoundaryFunc),
	)
}

func (b OptionBuilder) AsChrome() OptionBuilder {
	return mergeOptionBuilders(b, AsChrome())
}

// AsFirefox applies the Firefox browser profile (headers, TLS, and HTTP/2 behavior).
// @group Browser Profiles
//
// Applies to client configuration only.
// Example: use a Firefox profile
//
//	c := httpx.New(httpx.AsFirefox())
//	_ = c
func AsFirefox() OptionBuilder {
	return mergeOptionBuilders(
		TLSFingerprint(TLSFingerprintFirefoxKind),
		http2SettingsOption(firefoxHTTP2Settings...),
		http2ConnectionFlowOption(firefoxHTTP2ConnectionFlow),
		http2PriorityFramesOption(firefoxPriorityFrames...),
		pseudoHeaderOrder(firefoxPseudoHeaderOrder...),
		headerOrder(firefoxHeaderOrder...),
		Headers(firefoxHeaders),
		http2HeaderPriorityOption(firefoxHeaderPriority),
		multipartBoundary(firefoxMultipartBoundaryFunc),
	)
}

func (b OptionBuilder) AsFirefox() OptionBuilder {
	return mergeOptionBuilders(b, AsFirefox())
}

// AsSafari applies the Safari browser profile (headers, TLS, and HTTP/2 behavior).
// @group Browser Profiles
//
// Applies to client configuration only.
// Example: use a Safari profile
//
//	c := httpx.New(httpx.AsSafari())
//	_ = c
func AsSafari() OptionBuilder {
	return mergeOptionBuilders(
		TLSFingerprint(TLSFingerprintSafariKind),
		http2SettingsOption(safariHTTP2Settings...),
		http2ConnectionFlowOption(safariHTTP2ConnectionFlow),
		pseudoHeaderOrder(safariPseudoHeaderOrder...),
		headerOrder(safariHeaderOrder...),
		Headers(safariHeaders),
		http2HeaderPriorityOption(safariHeaderPriority),
		multipartBoundary(webkitMultipartBoundaryFunc),
	)
}

func (b OptionBuilder) AsSafari() OptionBuilder {
	return mergeOptionBuilders(b, AsSafari())
}

// AsMobile applies a mobile Chrome-like profile (headers, TLS, and HTTP/2 behavior).
// @group Browser Profiles
//
// Applies to client configuration only.
// Example: use a mobile profile
//
//	c := httpx.New(httpx.AsMobile())
//	_ = c
func AsMobile() OptionBuilder {
	return mergeOptionBuilders(
		TLSFingerprint(TLSFingerprintAndroidKind),
		http2SettingsOption(mobileHTTP2Settings...),
		http2ConnectionFlowOption(mobileHTTP2ConnectionFlow),
		pseudoHeaderOrder(mobilePseudoHeaderOrder...),
		headerOrder(mobileHeaderOrder...),
		Headers(mobileHeaders),
		http2HeaderPriorityOption(mobileHeaderPriority),
		multipartBoundary(webkitMultipartBoundaryFunc),
	)
}

func (b OptionBuilder) AsMobile() OptionBuilder {
	return mergeOptionBuilders(b, AsMobile())
}

func mergeOptionBuilders(builders ...OptionBuilder) OptionBuilder {
	var out OptionBuilder
	for _, b := range builders {
		out.ops = append(out.ops, b.ops...)
	}
	return out
}

func http2SettingsOption(settings ...http2.Setting) OptionBuilder {
	return OptionBuilder{}.add(clientOnly(func(c *Client) {
		if len(settings) == 0 {
			return
		}
		c.req.SetHTTP2SettingsFrame(settings...)
	}))
}

func http2ConnectionFlowOption(flow uint32) OptionBuilder {
	return OptionBuilder{}.add(clientOnly(func(c *Client) {
		c.req.SetHTTP2ConnectionFlow(flow)
	}))
}

func http2HeaderPriorityOption(priority http2.PriorityParam) OptionBuilder {
	return OptionBuilder{}.add(clientOnly(func(c *Client) {
		c.req.SetHTTP2HeaderPriority(priority)
	}))
}

func http2PriorityFramesOption(frames ...http2.PriorityFrame) OptionBuilder {
	return OptionBuilder{}.add(clientOnly(func(c *Client) {
		if len(frames) == 0 {
			return
		}
		c.req.SetHTTP2PriorityFrames(frames...)
	}))
}

// Identical for both Blink-based browsers (Chrome, Chromium, etc.) and WebKit-based browsers (Safari, etc.)
// Blink implementation: https://source.chromium.org/chromium/chromium/src/+/main:third_party/blink/renderer/platform/network/form_data_encoder.cc;drc=1d694679493c7b2f7b9df00e967b4f8699321093;l=130
// WebKit implementation: https://github.com/WebKit/WebKit/blob/47eea119fe9462721e5cc75527a4280c6d5f5214/Source/WebCore/platform/network/FormDataBuilder.cpp#L120
func webkitMultipartBoundaryFunc() string {
	const letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789AB"

	sb := strings.Builder{}
	sb.WriteString("----WebKitFormBoundary")

	for i := 0; i < 16; i++ {
		index, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters)-1)))
		if err != nil {
			panic(err)
		}

		sb.WriteByte(letters[index.Int64()])
	}

	return sb.String()
}

// Firefox implementation: https://searchfox.org/mozilla-central/source/dom/html/HTMLFormSubmission.cpp#355
func firefoxMultipartBoundaryFunc() string {
	sb := strings.Builder{}
	sb.WriteString("-------------------------")

	for i := 0; i < 3; i++ {
		var b [8]byte
		if _, err := io.ReadFull(rand.Reader, b[:]); err != nil {
			panic(err)
		}
		u32 := binary.LittleEndian.Uint32(b[:])
		s := strconv.FormatUint(uint64(u32), 10)

		sb.WriteString(s)
	}

	return sb.String()
}

var (
	chromeHTTP2Settings = []http2.Setting{
		{ID: http2.SettingHeaderTableSize, Val: 65536},
		{ID: http2.SettingEnablePush, Val: 0},
		{ID: http2.SettingMaxConcurrentStreams, Val: 1000},
		{ID: http2.SettingInitialWindowSize, Val: 6291456},
		{ID: http2.SettingMaxHeaderListSize, Val: 262144},
	}

	chromeHTTP2ConnectionFlow = uint32(15663105)

	chromePseudoHeaderOrder = []string{
		":method",
		":authority",
		":scheme",
		":path",
	}

	chromeHeaderOrder = []string{
		"host",
		"pragma",
		"cache-control",
		"sec-ch-ua",
		"sec-ch-ua-mobile",
		"sec-ch-ua-platform",
		"upgrade-insecure-requests",
		"user-agent",
		"accept",
		"sec-fetch-site",
		"sec-fetch-mode",
		"sec-fetch-user",
		"sec-fetch-dest",
		"referer",
		"accept-encoding",
		"accept-language",
		"cookie",
	}

	chromeHeaders = map[string]string{
		"pragma":                    "no-cache",
		"cache-control":             "no-cache",
		"sec-ch-ua":                 `"Not_A Brand";v="8", "Chromium";v="120", "Google Chrome";v="120"`,
		"sec-ch-ua-mobile":          "?0",
		"sec-ch-ua-platform":        `"macOS"`,
		"upgrade-insecure-requests": "1",
		"user-agent":                "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36",
		"accept":                    "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7",
		"sec-fetch-site":            "none",
		"sec-fetch-mode":            "navigate",
		"sec-fetch-user":            "?1",
		"sec-fetch-dest":            "document",
		"accept-language":           "zh-CN,zh;q=0.9",
	}

	chromeHeaderPriority = http2.PriorityParam{
		StreamDep: 0,
		Exclusive: true,
		Weight:    255,
	}

	firefoxHTTP2Settings = []http2.Setting{
		{ID: http2.SettingHeaderTableSize, Val: 65536},
		{ID: http2.SettingInitialWindowSize, Val: 131072},
		{ID: http2.SettingMaxFrameSize, Val: 16384},
	}

	firefoxHTTP2ConnectionFlow = uint32(12517377)

	firefoxPriorityFrames = []http2.PriorityFrame{
		{
			StreamID: 3,
			PriorityParam: http2.PriorityParam{
				StreamDep: 0,
				Exclusive: false,
				Weight:    200,
			},
		},
		{
			StreamID: 5,
			PriorityParam: http2.PriorityParam{
				StreamDep: 0,
				Exclusive: false,
				Weight:    100,
			},
		},
		{
			StreamID: 7,
			PriorityParam: http2.PriorityParam{
				StreamDep: 0,
				Exclusive: false,
				Weight:    0,
			},
		},
		{
			StreamID: 9,
			PriorityParam: http2.PriorityParam{
				StreamDep: 7,
				Exclusive: false,
				Weight:    0,
			},
		},
		{
			StreamID: 11,
			PriorityParam: http2.PriorityParam{
				StreamDep: 3,
				Exclusive: false,
				Weight:    0,
			},
		},
		{
			StreamID: 13,
			PriorityParam: http2.PriorityParam{
				StreamDep: 0,
				Exclusive: false,
				Weight:    240,
			},
		},
	}

	firefoxPseudoHeaderOrder = []string{
		":method",
		":path",
		":authority",
		":scheme",
	}

	firefoxHeaderOrder = []string{
		"user-agent",
		"accept",
		"accept-language",
		"accept-encoding",
		"referer",
		"cookie",
		"upgrade-insecure-requests",
		"sec-fetch-dest",
		"sec-fetch-mode",
		"sec-fetch-site",
		"sec-fetch-user",
		"te",
	}

	firefoxHeaders = map[string]string{
		"user-agent":                "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:120.0) Gecko/20100101 Firefox/120.0",
		"accept":                    "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8",
		"accept-language":           "zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2",
		"upgrade-insecure-requests": "1",
		"sec-fetch-dest":            "document",
		"sec-fetch-mode":            "navigate",
		"sec-fetch-site":            "same-origin",
		"sec-fetch-user":            "?1",
	}

	firefoxHeaderPriority = http2.PriorityParam{
		StreamDep: 13,
		Exclusive: false,
		Weight:    41,
	}

	safariHTTP2Settings = []http2.Setting{
		{ID: http2.SettingInitialWindowSize, Val: 4194304},
		{ID: http2.SettingMaxConcurrentStreams, Val: 100},
	}

	safariHTTP2ConnectionFlow = uint32(10485760)

	safariPseudoHeaderOrder = []string{
		":method",
		":scheme",
		":path",
		":authority",
	}

	safariHeaderOrder = []string{
		"accept",
		"sec-fetch-site",
		"cookie",
		"sec-fetch-dest",
		"accept-language",
		"sec-fetch-mode",
		"user-agent",
		"referer",
		"accept-encoding",
	}

	safariHeaders = map[string]string{
		"accept":          "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8",
		"sec-fetch-site":  "same-origin",
		"sec-fetch-dest":  "document",
		"accept-language": "zh-CN,zh-Hans;q=0.9",
		"sec-fetch-mode":  "navigate",
		"user-agent":      "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.6 Safari/605.1.15",
	}

	safariHeaderPriority = http2.PriorityParam{
		StreamDep: 0,
		Exclusive: false,
		Weight:    254,
	}

	mobileHTTP2Settings = chromeHTTP2Settings

	mobileHTTP2ConnectionFlow = chromeHTTP2ConnectionFlow

	mobilePseudoHeaderOrder = chromePseudoHeaderOrder

	mobileHeaderOrder = chromeHeaderOrder

	mobileHeaders = map[string]string{
		"pragma":                    "no-cache",
		"cache-control":             "no-cache",
		"sec-ch-ua":                 `"Not_A Brand";v="8", "Chromium";v="120", "Google Chrome";v="120"`,
		"sec-ch-ua-mobile":          "?1",
		"sec-ch-ua-platform":        `"Android"`,
		"upgrade-insecure-requests": "1",
		"user-agent":                "Mozilla/5.0 (Linux; Android 13; Pixel 7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Mobile Safari/537.36",
		"accept":                    "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7",
		"sec-fetch-site":            "none",
		"sec-fetch-mode":            "navigate",
		"sec-fetch-user":            "?1",
		"sec-fetch-dest":            "document",
		"accept-language":           "en-US,en;q=0.9",
	}

	mobileHeaderPriority = chromeHeaderPriority
)
