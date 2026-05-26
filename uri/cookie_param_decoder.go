package uri

import (
	"net/http"
)

type cookieParamDecoder struct {
	paramName string
	explode   bool
	req       *http.Request
}

func (d *cookieParamDecoder) DecodeValue() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// if NO error

func (d *cookieParamDecoder) DecodeArray(f func(Decoder) error) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *cookieParamDecoder) DecodeFields(f func(field string, d Decoder) error) error {
	_ = "STUB: not implemented"
	return nil
}
