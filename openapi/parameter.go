package openapi

import (
	"github.com/ogen-go/ogen/jsonschema"
	"github.com/ogen-go/ogen/location"
)

// ParameterLocation defines where OpenAPI parameter is located.
type ParameterLocation string

const (
	// LocationQuery is "query" parameter location.
	LocationQuery ParameterLocation = "query"
	// LocationHeader is "header" parameter location.
	LocationHeader ParameterLocation = "header"
	// LocationPath is "path" parameter location.
	LocationPath ParameterLocation = "path"
	// LocationCookie is "cookie" parameter location.
	LocationCookie ParameterLocation = "cookie"
)

// Query whether parameter location is query.
func (l ParameterLocation) Query() bool { _ = "STUB: not implemented"; return false }

// Header whether parameter location is header.
func (l ParameterLocation) Header() bool { _ = "STUB: not implemented"; return false }

// Path whether parameter location is path.
func (l ParameterLocation) Path() bool { _ = "STUB: not implemented"; return false }

// Cookie whether parameter location is cookie.
func (l ParameterLocation) Cookie() bool { _ = "STUB: not implemented"; return false }

// String implements fmt.Stringer.
func (l ParameterLocation) String() string {
	_ = "STUB: not implemented"

	// Parameter is an OpenAPI Operation Parameter.
	return ""
}

type Parameter struct {
	Ref Ref

	Name          string
	Description   string
	Deprecated    bool
	Schema        *jsonschema.Schema
	Content       *ParameterContent
	In            ParameterLocation
	Style         ParameterStyle
	Explode       bool
	Required      bool
	AllowReserved bool

	// XOgenName is the custom Go field name specified by x-ogen-name extension.
	XOgenName string

	location.Pointer `json:"-" yaml:"-"`
}

// ParameterContent describes OpenAPI Parameter content field.
//
// Parameter "content" field described as a map, and it MUST only contain one entry.
type ParameterContent struct {
	Name  string
	Media *MediaType
}

// ParameterStyle is an OpenAPI Parameter style.
type ParameterStyle string

// String implements fmt.Stringer.
func (s ParameterStyle) String() string {
	_ = "STUB: not implemented"

	// https://swagger.io/docs/specification/serialization/
	return ""
}

const (
	PathStyleSimple ParameterStyle = "simple"
	PathStyleLabel  ParameterStyle = "label"
	PathStyleMatrix ParameterStyle = "matrix"

	QueryStyleForm           ParameterStyle = "form"
	QueryStyleSpaceDelimited ParameterStyle = "spaceDelimited"
	QueryStylePipeDelimited  ParameterStyle = "pipeDelimited"
	QueryStyleDeepObject     ParameterStyle = "deepObject"

	HeaderStyleSimple ParameterStyle = "simple"

	CookieStyleForm ParameterStyle = "form"
)
