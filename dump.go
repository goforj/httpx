package httpx

import "github.com/goforj/godump"

// Dump prints values using the bundled godump formatter.
// @group Debugging
//
// Example: dump a response
//
//	res, err := httpx.Get[map[string]any](httpx.Default(), "https://httpbin.org/get")
//	_ = err
//	httpx.Dump(res)
func Dump(values ...any) {
	godump.Dump(values...)
}
