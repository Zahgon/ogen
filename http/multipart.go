package http

import (
	"io"
	"mime/multipart"
)

func randomBoundary() string { _ = "STUB: not implemented"; return "" }

// CreateMultipartBody is helper for streaming multipart/form-data.
func CreateMultipartBody(cb func(mw *multipart.Writer) error) (body io.ReadCloser, boundary string) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), ""
}
