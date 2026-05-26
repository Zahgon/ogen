package uri

import (
	"net/http"
)

type headerParamDecoder struct {
	paramName string
	explode   bool
	header    http.Header
}

func (d *headerParamDecoder) DecodeValue() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (d *headerParamDecoder) DecodeArray(f func(Decoder) error) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *headerParamDecoder) DecodeFields(f func(field string, d Decoder) error) error {
	_ = "STUB: not implemented"
	return nil
}
