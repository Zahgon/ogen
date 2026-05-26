// Package middleware provides a middleware interface for ogen.
package middleware

import (
	"context"
	"net/http"

	"github.com/ogen-go/ogen/openapi"
)

// ParameterKey is a map key for parameters.
type ParameterKey struct {
	// Name is the name of the parameter.
	Name string
	// In is the location of the parameter.
	In openapi.ParameterLocation
}

// Parameters is a map of parameters.
type Parameters map[ParameterKey]any

func (p Parameters) find(name string, in openapi.ParameterLocation) (v any, ok bool) {
	_ = "STUB: not implemented"
	return *new(any), false
}

// Query returns a parameter from the query.
func (p Parameters) Query(name string) (any, bool) {
	_ = "STUB: not implemented"
	return *new(any), false
}

// Header returns a parameter from the header.
func (p Parameters) Header(name string) (any, bool) {
	_ = "STUB: not implemented"
	return *new(any), false
}

// Path returns a parameter from the path.
func (p Parameters) Path(name string) (any, bool) {
	_ = "STUB: not implemented"
	return *new(any), false
}

// Cookie returns a parameter from the cookie.
func (p Parameters) Cookie(name string) (any, bool) {
	_ = "STUB: not implemented"
	return *new(any), false
}

// Request is request context type for middleware.
type Request struct {
	// Context is request context.
	Context context.Context
	// OperationName is the ogen operation name. It is guaranteed to be unique and not empty.
	OperationName string
	// OperationSummary is the ogen operation summary.
	OperationSummary string
	// OperationID is the spec operation ID, if any.
	OperationID string
	// Body is the operation request body. May be nil, if the operation has not body.
	Body any
	// RawBody is the operation request raw body. May be nil or empty, if the operation has no body.
	RawBody []byte
	// Params is the operation parameters.
	Params Parameters
	// Raw is the raw http request.
	Raw *http.Request
}

// SetContext sets Context in Request.
func (r *Request) SetContext(ctx context.Context) {
	_ = "STUB: not implemented"

	// Response is response type for middleware.
	return
}

type Response struct {
	// Type is the operation response type.
	Type any
}

type (
	// Next is the next middleware/handler in the chain.
	Next = func(req Request) (Response, error)
	// Middleware is middleware type.
	Middleware func(req Request, next Next) (Response, error)
)

// ChainMiddlewares chains middlewares into a single middleware, which will be executed in the order they are passed.
func ChainMiddlewares(m ...Middleware) Middleware {
	_ = "STUB: not implemented"
	return *new(Middleware)
}
