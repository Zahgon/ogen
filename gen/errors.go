package gen

import (
	"fmt"

	"github.com/go-faster/errors"

	"github.com/ogen-go/ogen/gen/ir"
)

type unimplementedError interface {
	unimplemented()
	error
}

var (
	_ = []interface {
		errors.Wrapper
		errors.Formatter
		fmt.Formatter
		error
	}{
		(*ErrParseSpec)(nil),
		(*ErrBuildRouter)(nil),
		(*ErrGoFormat)(nil),
	}
	_ = []interface {
		error
		unimplementedError
	}{
		(*ErrNotImplemented)(nil),
		(*ErrUnsupportedContentTypes)(nil),
	}
)

// ErrNotImplemented reports that feature is not implemented.
type ErrNotImplemented struct {
	Name string
}

func (e *ErrNotImplemented) unimplemented() {
	_ = "STUB: not implemented"

	// Error implements error.
	return
}

func (e *ErrNotImplemented) Error() string { _ = "STUB: not implemented"; return "" }

// ErrUnsupportedContentTypes reports that ogen does not support such content-type(s).
type ErrUnsupportedContentTypes struct {
	ContentTypes []string
}

func (e *ErrUnsupportedContentTypes) unimplemented() {
	_ = "STUB: not implemented"

	// Error implements error.
	return
}

func (e *ErrUnsupportedContentTypes) Error() string { _ = "STUB: not implemented"; return "" }

// ErrFieldsDiscriminatorInference reports fields discriminator inference failure.
type ErrFieldsDiscriminatorInference struct {
	Sum   *ir.Type
	Types []BadVariant
}

func (e *ErrFieldsDiscriminatorInference) unimplemented() {
	_ = "STUB: not implemented"

	// Error implements error.
	return
}

func (e *ErrFieldsDiscriminatorInference) Error() string { _ = "STUB: not implemented"; return "" }

// BadVariant describes a sum type variant for what we unable to infer discriminator.
type BadVariant struct {
	Type   *ir.Type
	Fields map[string][]*ir.Type
}

func (g *Generator) trySkip(err error, msg string, l position) error {
	_ = "STUB: not implemented"
	return nil
}

// Debug the original error "deep", to include the various messages added with Wrap*().

// Then log the brief error briefly with Info.

// Log the original error "deep", to include the various messages added with Wrap*().

func (g *Generator) fail(err error) error { _ = "STUB: not implemented"; return nil }

// ErrParseSpec reports that specification parsing failed.
type ErrParseSpec struct {
	err error
}

// Unwrap implements errors.Wrapper.
func (e *ErrParseSpec) Unwrap() error {
	_ = "STUB: not implemented"

	// FormatError implements errors.Formatter.
	return nil
}

func (e *ErrParseSpec) FormatError(p errors.Printer) (next error) {
	_ = "STUB: not implemented"
	return nil
}

// Format implements fmt.Formatter.
func (e *ErrParseSpec) Format(s fmt.State, verb rune) { _ = "STUB: not implemented"; return }

// Error implements error.
func (e *ErrParseSpec) Error() string { _ = "STUB: not implemented"; return "" }

// ErrBuildRouter reports that route tree building failed.
type ErrBuildRouter struct {
	err error
}

// Unwrap implements errors.Wrapper.
func (e *ErrBuildRouter) Unwrap() error {
	_ = "STUB: not implemented"

	// FormatError implements errors.Formatter.
	return nil
}

func (e *ErrBuildRouter) FormatError(p errors.Printer) (next error) {
	_ = "STUB: not implemented"
	return nil
}

// Format implements fmt.Formatter.
func (e *ErrBuildRouter) Format(s fmt.State, verb rune) { _ = "STUB: not implemented"; return }

// Error implements error.
func (e *ErrBuildRouter) Error() string { _ = "STUB: not implemented"; return "" }

// ErrGoFormat reports that generated code formatting failed.
type ErrGoFormat struct {
	err error
}

// Unwrap implements errors.Wrapper.
func (e *ErrGoFormat) Unwrap() error {
	_ = "STUB: not implemented"

	// FormatError implements errors.Formatter.
	return nil
}

func (e *ErrGoFormat) FormatError(p errors.Printer) (next error) {
	_ = "STUB: not implemented"
	return nil
}

// Format implements fmt.Formatter.
func (e *ErrGoFormat) Format(s fmt.State, verb rune) { _ = "STUB: not implemented"; return }

// Error implements error.
func (e *ErrGoFormat) Error() string { _ = "STUB: not implemented"; return "" }
