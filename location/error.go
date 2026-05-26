package location

import (
	"fmt"
	"io"

	"github.com/go-faster/errors"
)

var _ interface {
	errors.Wrapper
	errors.Formatter
	fmt.Formatter
	error
} = (*Error)(nil)

// Error is a wrapper for an error that has a location.
type Error struct {
	File File
	Pos  Position
	Err  error
}

// Unwrap implements errors.Wrapper.
func (e *Error) Unwrap() error {
	_ = "STUB: not implemented"

	// FormatError implements errors.Formatter.
	return nil
}

func (e *Error) FormatError(p errors.Printer) error { _ = "STUB: not implemented"; return nil }

// Format implements fmt.Formatter.
func (e *Error) Format(s fmt.State, verb rune) { _ = "STUB: not implemented"; return }

// Error implements error.
func (e *Error) Error() string { _ = "STUB: not implemented"; return "" }

// prettyPrint prints the error in a pretty way and returns true if it was printed successfully.
func (e *Error) prettyPrint(w io.Writer, opts PrintListingOptions) (handled bool, writeErr error) {
	_ = "STUB: not implemented"
	return false, nil
}

// Report is element of MultiError container.
type Report struct {
	File File
	Pos  Position
	Msg  string
}

// String returns textual represntation of Report.
func (r Report) String() string { _ = "STUB: not implemented"; return "" }

var _ interface {
	errors.Formatter
	fmt.Formatter
	error
} = (*MultiError)(nil)

// MultiError contains multiple Reports.
type MultiError struct {
	reports []Report
}

// Report adds report to the list.
func (e *MultiError) Report(file File, l Locator, msg string) { _ = "STUB: not implemented"; return }

// ReportPtr adds report to the list at given pointer.
func (e *MultiError) ReportPtr(ptr Pointer, msg string) { _ = "STUB: not implemented"; return }

func (e *MultiError) printSingle(printf func(format string, args ...any)) {
	_ = "STUB: not implemented"
	return
}

// FormatError implements errors.Formatter.
func (e *MultiError) FormatError(p errors.Printer) error { _ = "STUB: not implemented"; return nil }

// Format implements fmt.Formatter.
func (e *MultiError) Format(s fmt.State, verb rune) { _ = "STUB: not implemented"; return }

// Error implements error.
func (e *MultiError) Error() string { _ = "STUB: not implemented"; return "" }

const printLimit = 5

type reportChunk struct {
	Msg        string
	File       File
	Highlights []Highlight
}

func chunkReports(reports []Report, context int, hcolor ColorFunc) []reportChunk {
	_ = "STUB: not implemented"
	// Group Reports by Source (different files).
	return nil
}

// report with a non-empty message takes precedence.

// Line of previous position + its context + line in-between + second line context

// chunk with a non-empty message takes precedence.
// note the reverse order

// prettyPrint prints the error in a pretty way and returns true if it was printed successfully.
func (e *MultiError) prettyPrint(w io.Writer, opts PrintListingOptions) (handled bool, writeErr error) {
	_ = "STUB: not implemented"
	return false, nil
}

func printYAMLError(w io.Writer, err error, f File, opts PrintListingOptions) (handled bool, writeErr error) {
	_ = "STUB: not implemented"
	return false, nil
}

// Consider the error as handled if it is printed at least once.

// PrintPrettyError prints the error in a pretty way and returns true if it was printed successfully.
func PrintPrettyError(w io.Writer, color bool, err error) bool {
	_ = "STUB: not implemented"
	return false
}

// Set ColorFuncs to non-nil.

// TODO(tdakkota): handle write errors?
