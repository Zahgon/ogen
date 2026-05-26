package uri

import (
	"net/http"
)

type CookieDecoder struct {
	req *http.Request
}

func NewCookieDecoder(req *http.Request) *CookieDecoder { _ = "STUB: not implemented"; return nil }

type CookieParameterDecodingConfig struct {
	Name    string
	Explode bool
}

func (d *CookieDecoder) HasParam(cfg CookieParameterDecodingConfig) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *CookieDecoder) DecodeParam(cfg CookieParameterDecodingConfig, f func(Decoder) error) error {
	_ = "STUB: not implemented"
	return nil
}
