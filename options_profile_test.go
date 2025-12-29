package httpx

import (
	"reflect"
	"testing"

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

func getHTTP2Field(t *testing.T, c *Client, name string) reflect.Value {
	t.Helper()
	t2 := reflect.ValueOf(c.req.Transport).Elem().FieldByName("t2")
	if !t2.IsValid() || t2.IsNil() {
		t.Fatalf("expected http2 transport to be present")
	}
	return t2.Elem().FieldByName(name)
}
