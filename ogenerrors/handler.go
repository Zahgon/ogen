package ogenerrors

import (
	"context"
	"net/http"
)

// ErrorHandler is an error handler.
type ErrorHandler func(ctx context.Context, w http.ResponseWriter, r *http.Request, err error)

// ErrorCode returns HTTP code for given error.
//
// The default code is http.StatusInternalServerError.
func ErrorCode(err error) (code int) { _ = "STUB: not implemented"; return 0 }

// Takes precedence over Error.

// DefaultErrorHandler is the default error handler.
func DefaultErrorHandler(ctx context.Context, w http.ResponseWriter, r *http.Request, err error) {
	_ = "STUB: not implemented"
	return
}
