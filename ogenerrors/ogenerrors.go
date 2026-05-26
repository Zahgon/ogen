// Package ogenerrors contains ogen errors type definitions and helpers.
package ogenerrors

import (
	"fmt"

	"github.com/go-faster/errors"
)

// Error is an ogen error.
type Error interface {
	OperationName() string
	OperationID() string
	Code() int
	errors.Wrapper
	errors.Formatter
	fmt.Formatter
	error
}

var _ = []Error{
	new(SecurityError),
	new(DecodeParamsError),
	new(DecodeRequestError),
}

// OperationContext defines operation context for the error.
type OperationContext struct {
	Name string
	ID   string
}

// OperationName returns operation Name.
func (d OperationContext) OperationName() string {
	_ = "STUB: not implemented"

	// OperationID returns operation ID.
	return ""
}

func (d OperationContext) OperationID() string {
	_ = "STUB: not implemented"

	// SecurityError reports that error caused by security handler.
	return ""
}

type SecurityError struct {
	OperationContext
	Security string
	Err      error
}

// Code returns http code to respond.
func (d *SecurityError) Code() int { _ = "STUB: not implemented"; return 0 }

// Unwrap returns child error.
func (d *SecurityError) Unwrap() error {
	_ = "STUB: not implemented"

	// FormatError implements errors.Formatter.
	return nil
}

func (d *SecurityError) FormatError(p errors.Printer) (next error) {
	_ = "STUB: not implemented"
	return nil
}

// Format implements fmt.Formatter.
func (d *SecurityError) Format(s fmt.State, verb rune) { _ = "STUB: not implemented"; return }

// Error implements error.
func (d *SecurityError) Error() string { _ = "STUB: not implemented"; return "" }

// DecodeRequestError reports that error caused by request decoder.
type DecodeRequestError struct {
	OperationContext
	Err error
}

// Code returns http code to respond.
func (d *DecodeRequestError) Code() int { _ = "STUB: not implemented"; return 0 }

// Unwrap returns child error.
func (d *DecodeRequestError) Unwrap() error {
	_ = "STUB: not implemented"

	// FormatError implements errors.Formatter.
	return nil
}

func (d *DecodeRequestError) FormatError(p errors.Printer) (next error) {
	_ = "STUB: not implemented"
	return nil
}

// Format implements fmt.Formatter.
func (d *DecodeRequestError) Format(s fmt.State, verb rune) { _ = "STUB: not implemented"; return }

// Error implements error.
func (d *DecodeRequestError) Error() string { _ = "STUB: not implemented"; return "" }
