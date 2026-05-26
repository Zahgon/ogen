package gen

import (
	"strings"

	"github.com/ogen-go/ogen/gen/ir"
)

// generateUniqueValidators generates validateUnique[TypeName]() functions
// for arrays of complex types that require hash-based duplicate detection.
func (g *Generator) generateUniqueValidators(fs FileSystem, pkgName string) error {
	_ = "STUB: not implemented"
	return nil
}

// File header

// Generate function for each type

// writeValidateUnique generates a single validateUnique[TypeName]() function
func writeValidateUnique(b *strings.Builder, spec *ir.EqualityMethodSpec) {
	_ = "STUB: not implemented"
	return
}

// Depth limit panic recovery

// Hash bucket structure

// Duplicate detection loop
