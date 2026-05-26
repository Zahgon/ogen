package ogenerrors

import (
	"fmt"

	"github.com/go-faster/errors"

	"github.com/ogen-go/ogen/openapi"
)

// DecodeParamsError reports that error caused by params decoder.
type DecodeParamsError struct {
	OperationContext
	Err error
}

// Code returns http code to respond.
func (d *DecodeParamsError) Code() int { _ = "STUB: not implemented"; return 0 }

// Unwrap returns child error.
func (d *DecodeParamsError) Unwrap() error {
	_ = "STUB: not implemented"

	// FormatError implements errors.Formatter.
	return nil
}

func (d *DecodeParamsError) FormatError(p errors.Printer) (next error) {
	_ = "STUB: not implemented"
	return nil
}

// Format implements fmt.Formatter.
func (d *DecodeParamsError) Format(s fmt.State, verb rune) { _ = "STUB: not implemented"; return }

// Error implements error.
func (d *DecodeParamsError) Error() string { _ = "STUB: not implemented"; return "" }

// DecodeParamError reports that error caused by parameter decoder.
type DecodeParamError struct {
	Name string
	In   openapi.ParameterLocation
	Err  error
}

// Unwrap returns child error.
func (d *DecodeParamError) Unwrap() error {
	_ = "STUB: not implemented"

	// FormatError implements errors.Formatter.
	return nil
}

func (d *DecodeParamError) FormatError(p errors.Printer) (next error) {
	_ = "STUB: not implemented"
	return nil
}

// Format implements fmt.Formatter.
func (d *DecodeParamError) Format(s fmt.State, verb rune) { _ = "STUB: not implemented"; return }

// Error implements error.
func (d *DecodeParamError) Error() string { _ = "STUB: not implemented"; return "" }
