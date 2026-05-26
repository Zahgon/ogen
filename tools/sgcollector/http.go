package main

import (
	"net/http"
)

type filterTransport struct {
	next    http.RoundTripper
	allowed map[string]struct{}
}

func (f filterTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
