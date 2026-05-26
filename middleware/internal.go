package middleware

import "context"

// HookMiddleware is a helper that does ogen request type -> Request type conversion.
//
// NB: this is an internal func, not intended for public use.
func HookMiddleware[RequestType, ParamsType, ResponseType any](
	m Middleware,
	req Request,
	unpack func(Parameters) ParamsType,
	cb func(context.Context, RequestType, ParamsType) (ResponseType, error),
) (r ResponseType, err error) {
	_ = "STUB: not implemented"
	return *new(ResponseType), nil
}
