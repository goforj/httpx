package httpx

import "github.com/imroc/req/v3"

// OutputFile streams the response body to a file path.
// @group Download Options
//
// Applies to individual requests only.
// Example: download to file
//
//	c := httpx.New()
//	res, err := httpx.Get[map[string]any](c, "https://httpbin.org/bytes/1024", httpx.OutputFile("/tmp/file.bin"))
//	_ = err
//	httpx.Dump(res) // dumps map[string]any
//	// #map[string]interface {} {}
func OutputFile(path string) OptionBuilder {
	return OptionBuilder{}.OutputFile(path)
}

func (b OptionBuilder) OutputFile(path string) OptionBuilder {
	return b.add(requestOnly(func(r *req.Request) {
		r.SetOutputFile(path)
	}))
}
