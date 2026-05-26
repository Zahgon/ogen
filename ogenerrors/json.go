package ogenerrors

import (
	"fmt"

	"github.com/go-faster/errors"
)

var _ interface {
	errors.Wrapper
	errors.Formatter
	fmt.Formatter
	error
} = (*DecodeBodyError)(nil)

// DecodeBodyError occurs when request or response body cannot be decoded.
type DecodeBodyError struct {
	ContentType string
	Body        []byte
	Err         error
}

// Unwrap returns child error.
func (d *DecodeBodyError) Unwrap() error {
	_ = "STUB: not implemented"

	// FormatError implements errors.Formatter.
	return nil
}

func (d *DecodeBodyError) FormatError(p errors.Printer) (next error) {
	_ = "STUB: not implemented"
	return nil
}

// Format implements fmt.Formatter.
func (d *DecodeBodyError) Format(s fmt.State, verb rune) { _ = "STUB: not implemented"; return }

// Error implements error.
func (d *DecodeBodyError) Error() string { _ = "STUB: not implemented"; return "" }
