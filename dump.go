package httpx

import "github.com/goforj/godump"

var dumpDumper = godump.NewDumper(godump.WithoutHeader())

// Dump prints values using the bundled godump formatter.
// @group Debugging
//
// Example: dump a response
//
//	res, err := httpx.Get[map[string]any](httpx.Default(), "https://httpbin.org/uuid")
//	_ = err
//	httpx.Dump(res)
//	// #map[string]interface {} {
//	//   uuid => "<uuid>" #string
//	// }
func Dump(values ...any) {
	dumpDumper.Dump(values...)
}
