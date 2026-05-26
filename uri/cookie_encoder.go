package uri

import (
	"net/http"
)

type CookieEncoder struct {
	req *http.Request
}

func NewCookieEncoder(req *http.Request) *CookieEncoder { _ = "STUB: not implemented"; return nil }

type CookieParameterEncodingConfig struct {
	Name    string
	Explode bool
}

func (e *CookieEncoder) EncodeParam(cfg CookieParameterEncodingConfig, f func(Encoder) error) error {
	_ = "STUB: not implemented"
	return nil
}
