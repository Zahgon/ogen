package openapi

import (
	"github.com/ogen-go/ogen/location"
)

// Operation is an OpenAPI Operation.
type Operation struct {
	Tags        []string // optional
	OperationID string   // optional
	Summary     string   // optional
	Description string   // optional
	Deprecated  bool     // optional

	HTTPMethod  string
	Path        Path
	Parameters  []*Parameter
	RequestBody *RequestBody // optional

	// Security requirements.
	Security SecurityRequirements

	// Operation responses.
	Responses Responses

	XOgenOperationGroup string // Extension field for operation grouping.

	location.Pointer `json:"-" yaml:"-"`
}

// RequestBody of an OpenAPI Operation.
type RequestBody struct {
	Ref         Ref
	Description string
	Required    bool
	Content     map[string]*MediaType

	location.Pointer `json:"-" yaml:"-"`
}

// Header is an OpenAPI Header definition.
type Header = Parameter

// Response is an OpenAPI Response definition.
type Response struct {
	Ref         Ref
	Description string
	Headers     map[string]*Header
	Content     map[string]*MediaType
	// Links map[string]*Link

	location.Pointer `json:"-" yaml:"-"`
}

// Responses contains a list of parsed OpenAPI Responses.
type Responses struct {
	StatusCode map[int]*Response
	Pattern    [5]*Response
	Default    *Response

	location.Pointer `json:"-" yaml:"-"`
}

// Add adds a response to the Responses.
func (r *Responses) Add(pattern string, resp *Response) error {
	_ = "STUB: not implemented"
	return nil
}

// Do not return parsing error, it could be a bit confusing.
