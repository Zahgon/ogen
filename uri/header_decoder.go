package uri

import (
	"net/http"
)

type HeaderDecoder struct {
	header http.Header
}

func NewHeaderDecoder(header http.Header) *HeaderDecoder { _ = "STUB: not implemented"; return nil }

type HeaderParameterDecodingConfig struct {
	Name    string
	Explode bool
}

func (d *HeaderDecoder) HasParam(cfg HeaderParameterDecodingConfig) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *HeaderDecoder) DecodeParam(cfg HeaderParameterDecodingConfig, f func(Decoder) error) error {
	_ = "STUB: not implemented"
	return nil
}
