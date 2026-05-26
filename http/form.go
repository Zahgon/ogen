package http

import (
	"net/http"
	"net/url"
)

// ParseForm is optimized version of http.Request.ParseForm.
//
// Difference from http.Request.ParseForm:
//   - This function does not modify any fields of http.Request. The only copy of the form values is returned.
//   - This function does not check Content-Type header.
func ParseForm(r *http.Request) (url.Values, error) {
	_ = "STUB: not implemented"
	return *new(url.Values), nil
}

// TODO(tdakkota): implement streaming parser?
