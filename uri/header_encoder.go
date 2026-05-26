package uri

import "net/http"

type HeaderEncoder struct {
	header http.Header
}

func NewHeaderEncoder(header http.Header) *HeaderEncoder { _ = "STUB: not implemented"; return nil }

type HeaderParameterEncodingConfig struct {
	Name    string
	Explode bool
}

func (e *HeaderEncoder) EncodeParam(cfg HeaderParameterEncodingConfig, f func(Encoder) error) error {
	_ = "STUB: not implemented"
	return nil
}

// FIXME(tdakkota): probable we should return the error during encoding

func (e *HeaderEncoder) Header() http.Header { _ = "STUB: not implemented"; return *new(http.Header) }
