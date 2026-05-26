package location

import (
	"fmt"
	"io"

	"golang.org/x/exp/constraints"
)

// ColorFunc defines a simple printer callback.
type ColorFunc func(w io.Writer, s string, args ...any) (int, error)

// PrintListingOptions is a set of options for PrintListing.
type PrintListingOptions struct {
	// Context is the number of lines to print before and after the error line.
	//
	// If is zero, the default value 5 is used.
	Context int
	// MsgColor sets message color.
	MsgColor ColorFunc
	// TextColor sets text color.
	PlainColor ColorFunc
}

// WithoutColor creates a copy of the options with disabled color.
func (o PrintListingOptions) WithoutColor() PrintListingOptions {
	_ = "STUB: not implemented"
	return *new(PrintListingOptions)
}

const defaultContext = 3

func (o *PrintListingOptions) setDefaults() { _ = "STUB: not implemented"; return }

const (
	// BugLine is a fallback line when the line is not available.
	BugLine = `Cannot render line properly, please fill a bug report`

	leftPad           = "  "
	verticalBorder    = "|"
	horizontalPointer = "\u2192"
)

// PrintListing prints given message with line number and file listing to the writer.
//
// The context parameter defines the number of lines to print before and after.
func (f File) PrintListing(w io.Writer, msg string, pos Position, opts PrintListingOptions) error {
	_ = "STUB: not implemented"
	return nil
}

// Highlight is a highlighted position.
type Highlight struct {
	Pos   Position
	Color ColorFunc
}

// clamp keeps val in given boundaries.
func clamp[T constraints.Integer](val, lo, hi T) T { _ = "STUB: not implemented"; return *new(T) }

func log10(val int) (r int) { _ = "STUB: not implemented"; return 0 }

type lineNumberPad struct {
	pad, line int
}

func (p lineNumberPad) Format(f fmt.State, verb rune) { _ = "STUB: not implemented"; return }

// PrintHighlights prints all given highlights.
func (f File) PrintHighlights(w io.Writer, msg string, highlights []Highlight, opts PrintListingOptions) error {
	_ = "STUB: not implemented"
	return nil
}

// Line starts from 1, but index starts from 0.

// Line number is 1-based.

// Line number is 1-based, but index is 0-based.

// Print lines.

// TODO(tdakkota): column pointer?
