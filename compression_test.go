package httpx

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/klauspost/compress/zstd"
)

// TestAutomaticZstdResponseDecompression preserves opt-in decoding through the underlying HTTP client.
func TestAutomaticZstdResponseDecompression(t *testing.T) {
	encoder, err := zstd.NewWriter(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer encoder.Close()

	const payload = "a compressed HTTP response"
	compressed := encoder.EncodeAll([]byte(payload), nil)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Encoding", "zstd")
		_, _ = w.Write(compressed)
	}))
	t.Cleanup(server.Close)

	client := New()
	client.Raw().EnableAutoDecompress()
	got, err := client.Get[string](server.URL)
	if err != nil {
		t.Fatal(err)
	}
	if got != payload {
		t.Fatalf("decoded response = %q, want %q", got, payload)
	}
}
