package uri

import (
	"net/url"
)

type QueryDecoder struct {
	values url.Values
}

func NewQueryDecoder(values url.Values) *QueryDecoder { _ = "STUB: not implemented"; return nil }

type QueryParameterDecodingConfig struct {
	Name    string
	Style   QueryStyle
	Explode bool
	Fields  []QueryParameterObjectField // Only for object param.
}

type QueryParameterObjectField struct {
	Name     string
	Required bool
}

func (d *QueryDecoder) HasParam(cfg QueryParameterDecodingConfig) error {
	_ = "STUB: not implemented"
	return nil

	// https://swagger.io/docs/specification/serialization/
}

// For deepObject with no predefined fields (additionalProperties/maps),
// check for keys matching the "paramName[" prefix pattern.

func (d *QueryDecoder) DecodeParam(cfg QueryParameterDecodingConfig, f func(Decoder) error) error {
	_ = "STUB: not implemented"
	return nil
}
