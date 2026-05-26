// Package http implements crazy ideas for http optimizations that should be
// mostly std compatible.
package http

import (
	"context"
	"io"
	"net/http"
	"net/url"

	"golang.org/x/sync/errgroup"
)

// Client represents http client.
type Client interface {
	Do(r *http.Request) (*http.Response, error)
}

// NewRequest creates a new http.Request.
func NewRequest(ctx context.Context, method string, u *url.URL) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func initRequest(req *http.Request, contentType string) { _ = "STUB: not implemented"; return }

// SetBody sets request body.
func SetBody(req *http.Request, body io.Reader, contentType string) {
	_ = "STUB: not implemented"
	return
}

// SetCloserBody sets request body which should be closed after request.
func SetCloserBody(req *http.Request, body io.ReadCloser, contentType string) {
	_ = "STUB: not implemented"
	return
}

// CreateBodyWriter is a helper to create a reader from a writer body.
func CreateBodyWriter(cb func(w io.Writer) error) io.ReadCloser {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser)
}

type bodyReader struct {
	r  *io.PipeReader
	w  *io.PipeWriter
	wg *errgroup.Group
}

func (w bodyReader) Read(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (w bodyReader) Close() (rerr error) { _ = "STUB: not implemented"; return nil }
