package ir

import (
	"github.com/ogen-go/ogen/openapi"
)

// WebhookInfo contains information about webhook.
type WebhookInfo struct {
	// Name is the name of the webhook.
	Name string
}

type Operation struct {
	Name           string
	Summary        string
	Description    string
	Deprecated     bool
	WebhookInfo    *WebhookInfo
	PathParts      []*PathPart
	Params         []*Parameter
	Request        *Request
	Responses      *Responses
	Security       SecurityRequirements
	Spec           *openapi.Operation
	OperationGroup string
}

type OperationGroup struct {
	Name       string
	Operations []*Operation
}

// OTELAttribute represents OpenTelemetry attribute defined by otelogen package.
type OTELAttribute struct {
	// Key is a name of the attribute constructor in otelogen package.
	Key string
	// Value is a value of the attribute.
	Value string
}

// String returns call to the constructor of this attribute.
func (a OTELAttribute) String() string { _ = "STUB: not implemented"; return "" }

// OTELAttributes returns OpenTelemetry attributes for this operation.
func (op Operation) OTELAttributes() (r []OTELAttribute) { _ = "STUB: not implemented"; return nil }

func (op Operation) PrettyOperationID() string { _ = "STUB: not implemented"; return "" }

func (op Operation) GoDoc() []string { _ = "STUB: not implemented"; return nil }

// HasRawResponse returns true if the operation has any response content types
// marked with x-ogen-raw-response: true.
func (op Operation) HasRawResponse() bool { _ = "STUB: not implemented"; return false }

// Check all responses

type PathPart struct {
	Raw   string
	Param *Parameter
}

func (p PathPart) String() string { _ = "STUB: not implemented"; return "" }

type Parameter struct {
	Name string
	Type *Type
	Spec *openapi.Parameter
	Tag  Tag
}

func (op Parameter) GoDoc() []string { _ = "STUB: not implemented"; return nil }

// Default returns default value of this field, if it is set.
func (op Parameter) Default() Default { _ = "STUB: not implemented"; return *new(Default) }

func doTakePtr(t *Type) bool { _ = "STUB: not implemented"; return false }

func reqRespGoType(t *Type) string { _ = "STUB: not implemented"; return "" }

type Request struct {
	Type      *Type
	EmptyBody *Type
	Contents  map[ContentType]Media
	Spec      *openapi.RequestBody
}

// DoTakePtr returns true if type should be taken by pointer.
func (r *Request) DoTakePtr() bool { _ = "STUB: not implemented"; return false }

// GoType returns Go type of this response.
func (r *Request) GoType() string { _ = "STUB: not implemented"; return "" }

type Responses struct {
	Type       *Type
	Pattern    [5]*Response
	StatusCode map[int]*Response
	Default    *Response
}

// DoTakePtr returns true if type should be taken by pointer.
func (r *Responses) DoTakePtr() bool { _ = "STUB: not implemented"; return false }

// DoPass whether response type should be present in result tuple.
func (r *Responses) DoPass() bool {
	_ = "STUB: not implemented"
	// In case of pattern responses or default response, the response type
	// has a StatusCode field, so it should be passed.
	return false
}

// Do not pass response type if it is empty struct.

// GoType returns Go type of this response.
func (r *Responses) GoType() string { _ = "STUB: not implemented"; return "" }

// ResultTuple returns result tuple for this response.
func (r *Responses) ResultTuple(a, b string) string { _ = "STUB: not implemented"; return "" }

// Ensure that all result tuple elements are named
// if any of them already is.

func (r *Responses) HasPattern() bool { _ = "STUB: not implemented"; return false }

type Response struct {
	NoContent *Type
	Contents  map[ContentType]Media
	Spec      *openapi.Response
	Headers   map[string]*Parameter

	// Indicates that all response types
	// are wrappers with StatusCode field.
	WithStatusCode bool

	// Indicates that all response types
	// are wrappers with response header fields.
	WithHeaders bool

	// Note that if NoContent is false
	// (i.e. response has specified contents)
	// and (WithStatusCode || WithHeaders) == true
	// all wrapper types will also have a Response field
	// which will contain the actual response body.
}

func (s Response) ResponseInfo(otel bool) []ResponseInfo { _ = "STUB: not implemented"; return nil }
