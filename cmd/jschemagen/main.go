package main

import (
	"fmt"
	"os"

	"github.com/ogen-go/ogen/jsonschema"
)

// StringArrayFlag is a string array flag.
type StringArrayFlag []string

// String implements fmt.Stringer.
func (i *StringArrayFlag) String() string { _ = "STUB: not implemented"; return "" }

// Set implements flag.Value.
func (i *StringArrayFlag) Set(value string) error { _ = "STUB: not implemented"; return nil }

func inferFileName(
	targetFile *string,
	typeName string,
	rawSchema jsonschema.RawSchema,
	trimPrefixes StringArrayFlag,
) {
	_ = "STUB: not implemented"
	// Output file already set.
	return
}

// Check that type name contains only valid path characters.

func run() error { _ = "STUB: not implemented"; return nil }

func main() {
	if err := run(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%+v\n", err)
		os.Exit(1)
	}
}
