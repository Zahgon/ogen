package uri

import (
	"net/url"
)

type QueryStyle string

const (
	QueryStyleForm           QueryStyle = "form"
	QueryStyleSpaceDelimited QueryStyle = "spaceDelimited"
	QueryStylePipeDelimited  QueryStyle = "pipeDelimited"
	QueryStyleDeepObject     QueryStyle = "deepObject"
)

type queryParamEncoder struct {
	*receiver
	values url.Values

	paramName string     // immutable
	style     QueryStyle // immutable
	explode   bool       // immutable
}

func (e *queryParamEncoder) serialize() error { _ = "STUB: not implemented"; return nil }

func (e *queryParamEncoder) encodeValue() error { _ = "STUB: not implemented"; return nil }

func (e *queryParamEncoder) encodeArray() error { _ = "STUB: not implemented"; return nil }

func (e *queryParamEncoder) encodeObject() error { _ = "STUB: not implemented"; return nil }
