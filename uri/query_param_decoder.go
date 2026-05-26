package uri

import (
	"net/url"
)

type queryParamDecoder struct {
	values       url.Values
	objectFields []QueryParameterObjectField

	paramName string
	style     QueryStyle // immutable
	explode   bool       // immutable
}

func (d *queryParamDecoder) DecodeValue() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (d *queryParamDecoder) DecodeArray(f func(d Decoder) error) error {
	_ = "STUB: not implemented"
	return nil
}

// do not decode `?param=` as `[""]` and leave the parameter as whatever zero value it has

func (d *queryParamDecoder) DecodeFields(f func(name string, d Decoder) error) error {
	_ = "STUB: not implemented"
	return nil
}

// Maintain a map of seen keys to avoid additional processing.

// With PatternProperties/AdditionalProperties, we do not know the names of these properties.
// Therefore, we just iterate over all query parameters that start with the prefix.

// This query parameter does not start with the prefix, so we skip it.

// To construct the field name, we need to remove the prefix while also removing the brackets.
// For example:
//   ?paramName[fieldName]=value => fieldName=value
//   ?paramName[fieldName][subField1]=value => fieldName[subField1]=value
//   ?paramName[fieldName][subField1][subSubField1]=value => fieldName[subField1][subSubField1]=value

// Remove the first open and close brackets to get the field name.

// This value was well-formed as part of the predefined d.objectFields, so we do not need to
// process it again.
