package main

import (
	"context"
	"net/http"
	"net/url"
	"regexp"

	"github.com/go-faster/errors"

	"github.com/ogen-go/ogen/location"
)

var errPanic = errors.New("panic")

func validateJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

var gitPathRegex = regexp.MustCompile(`^/(?P<owner>[^/]+)/(?P<repo>[^/]+)/(-/blob|blob)/(?P<ref>[^/]+)/(?P<path>.*)$`)

func getRootURL(m FileMatch) (*url.URL, bool) { _ = "STUB: not implemented"; return nil, false }

func worker(ctx context.Context, m FileMatch, r *Reporters, skipWrite bool) (rErr error) {
	_ = "STUB: not implemented"
	return nil
}

type nopFs struct{}

func (n nopFs) WriteFile(string, []byte) error { _ = "STUB: not implemented"; return nil }

func workerHTTPClient() *http.Client { _ = "STUB: not implemented"; return nil }

type file struct {
	data    []byte
	isYAML  bool
	name    string
	rootURL *url.URL
}

func (f file) location() location.File { _ = "STUB: not implemented"; return *new(location.File) }

func generate(f file, skipWrite bool) *GenerateError { _ = "STUB: not implemented"; return nil }
