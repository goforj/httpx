package httpx

import (
	"fmt"
	"net/http"

	"github.com/imroc/req/v3"
)

type HTTPError struct {
	StatusCode int
	Status     string
	Body       []byte
	Header     http.Header
}

// Error returns a short, human-friendly summary of the HTTP error.
//
// Example: check for HTTP errors
//
//	type User struct {
//		Name string `json:"name"`
//	}
//
//	c := httpx.New()
//	res := httpx.Get[User](c, "https://example.com/users/1")
//	var httpErr *httpx.HTTPError
//	if errors.As(res.Err, &httpErr) {
//		_ = httpErr.StatusCode
//	}
func (e *HTTPError) Error() string {
	return fmt.Sprintf("httpx: http %d %s", e.StatusCode, e.Status)
}

func newHTTPError(resp *req.Response) *HTTPError {
	if resp == nil {
		return &HTTPError{StatusCode: 0, Status: "missing response"}
	}
	return &HTTPError{
		StatusCode: resp.StatusCode,
		Status:     resp.Status,
		Body:       resp.Bytes(),
		Header:     resp.Header,
	}
}
