package validate

import (
	"net/http"

	"github.com/ogen-go/ogen/ogenregex"

	"github.com/go-faster/errors"
)

// ErrFieldRequired reports that a field is required, but not found.
var ErrFieldRequired = errors.New("field required")

// Error represents validation error.
type Error struct {
	Fields []FieldError
}

// Error implements error.
func (e *Error) Error() string { _ = "STUB: not implemented"; return "" }

// FieldError is failed validation on field.
type FieldError struct {
	Name  string
	Error error
}

// ErrBodyRequired reports that request body is required but server got empty request.
var ErrBodyRequired = errors.New("body required")

// InvalidContentTypeError reports that decoder got unexpected content type.
type InvalidContentTypeError struct {
	ContentType string
}

// InvalidContentTypeError implements error.
func (e *InvalidContentTypeError) Error() string { _ = "STUB: not implemented"; return "" }

// InvalidContentType creates new InvalidContentTypeError.
func InvalidContentType(contentType string) error { _ = "STUB: not implemented"; return nil }

// UnexpectedStatusCodeError reports that client got unexpected status code.
type UnexpectedStatusCodeError struct {
	StatusCode int
	Payload    *http.Response
}

// UnexpectedStatusCodeWithResponse creates new UnexpectedStatusCode.
func UnexpectedStatusCodeWithResponse(response *http.Response) error {
	_ = "STUB: not implemented"
	return nil
}

// response.Body is defer-closed by caller.
// we want to retain it in Payload.

// UnexpectedStatusCode creates new UnexpectedStatusCode.
//
// Deprecated: client codes generated a while ago used this function.
// Kept here solely for backward compatibility to them.
func UnexpectedStatusCode(statusCode int) error { _ = "STUB: not implemented"; return nil }

// UnexpectedStatusCodeError implements error.
func (e *UnexpectedStatusCodeError) Error() string { _ = "STUB: not implemented"; return "" }

// ErrNilPointer reports that use Validate, but receiver pointer is nil.
var ErrNilPointer = errors.New("nil pointer")

// MinLengthError reports that len less than minimum.
type MinLengthError struct {
	Len       int
	MinLength int
}

// MinLengthError implements error.
func (e *MinLengthError) Error() string { _ = "STUB: not implemented"; return "" }

// MaxLengthError reports that len greater than maximum.
type MaxLengthError struct {
	Len       int
	MaxLength int
}

// MaxLengthError implements error.
func (e *MaxLengthError) Error() string { _ = "STUB: not implemented"; return "" }

// NoRegexMatchError reports that value have no regexp match.
type NoRegexMatchError struct {
	Pattern ogenregex.Regexp
}

// MaxLengthError implements error.
func (e *NoRegexMatchError) Error() string { _ = "STUB: not implemented"; return "" }

// DuplicateItemsError indicates duplicate items in a uniqueItems array.
type DuplicateItemsError struct {
	// Indices contains all indices where duplicates were found.
	// First element is the original, subsequent are duplicates.
	Indices []int
}

// Error implements error.
func (e *DuplicateItemsError) Error() string { _ = "STUB: not implemented"; return "" }

// DepthLimitError indicates nesting depth limit was exceeded.
type DepthLimitError struct {
	// MaxDepth is the configured maximum depth.
	MaxDepth int

	// TypeName is the type being compared when limit was hit.
	TypeName string
}

// Error implements error.
func (e *DepthLimitError) Error() string { _ = "STUB: not implemented"; return "" }
