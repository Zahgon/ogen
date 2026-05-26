package gen

import (
	"strings"

	"github.com/ogen-go/ogen/gen/ir"
)

// generateEqualityMethodsWithFS generates Equal() and Hash() methods using a FileSystem
func (g *Generator) generateEqualityMethodsWithFS(fs FileSystem, pkgName string) error {
	_ = "STUB: not implemented"
	return nil
}

// generateEqualMethod generates an Equal() method for a type
func (g *Generator) generateEqualMethod(spec *ir.EqualityMethodSpec, pkgName string, fs FileSystem) error {
	_ = "STUB: not implemented"
	return nil

	// File header comment
}

// Method signature

// Depth limit check

// Generate field comparisons

// Write to file

// writeFieldComparison generates comparison code for a single field
func (g *Generator) writeFieldComparison(b *strings.Builder, field ir.FieldEqualitySpec, hasDepth bool) {
	_ = "STUB: not implemented"
	return
}

// Byte slices need bytes.Equal()

// Simple equality check

// Optional field comparison (OptT types)

// Optional wrapper around nested object - call Equal()

// Optional wrapper around map - need custom comparison

// Optional wrapper around array - need iteration

// Optional wrapper around primitive - use !=

// Nullable field comparison (NilT types)

// Nullable wrapper around nested object - call Equal()

// Nullable wrapper around map - need custom comparison

// Nullable wrapper around array - need iteration

// Nullable wrapper around primitive - use !=

// Pointer field comparison

// Array field comparison

// Check if array elements are nullable wrappers

// Nullable wrapper elements - compare Null flag then Value

// Check if the wrapped value is a struct (needs Equal()) or primitive (use !=)

// Value is a struct - call Equal()

// Value is a primitive - use !=

// Struct elements - call Equal()

// Primitive elements can use !=

// Map field comparison

// Nested object comparison (recursive Equal() call)

// generateHashMethod generates a Hash() method for a type
func (g *Generator) generateHashMethod(spec *ir.EqualityMethodSpec, pkgName string, fs FileSystem) error {
	_ = "STUB: not implemented"
	return nil

	// File header comment
}

// Imports

// Method signature

// Initialize hash

// Hash each field

// Convert FieldEqualitySpec to FieldHashSpec

// Write to file

// writeFieldHash generates hashing code for a single field
func (g *Generator) writeFieldHash(b *strings.Builder, field ir.FieldHashSpec) {
	_ = "STUB: not implemented"
	return
}

// Hash primitive field

// Hash optional field (presence + value)

// Optional wrapper around nested object - call Hash()

// Optional wrapper around primitive - format value

// Hash nullable field

// Nullable wrapper around nested object - call Hash()

// Nullable wrapper around primitive - format value

// Hash pointer field

// Hash array field

// Hash map field (sorted keys for deterministic hash)

// Hash nested object (incorporate its Hash() result)

// writeFormattedCode formats and writes code to a file using FileSystem
func writeFormattedCode(fs FileSystem, filename string, content []byte) error {
	_ = "STUB: not implemented"
	// Format with goimports
	return nil
}

// If formatting fails, return error (don't write malformed code)

// Write to file using FileSystem interface

// Fallback to os.WriteFile for testing

// snakeCase converts PascalCase to snake_case
func snakeCase(s string) string { _ = "STUB: not implemented"; return "" }

// writeMapComparison generates map comparison code
func (g *Generator) writeMapComparison(b *strings.Builder, aMap, bMap, indent string) {
	_ = "STUB: not implemented"
	return
}

// writeArrayComparison generates array comparison code
func (g *Generator) writeArrayComparison(b *strings.Builder, aArray, bArray, indent string, isArrayOfStructs, hasDepth bool) {
	_ = "STUB: not implemented"
	return
}

// writeArrayComparisonWithNullable generates array comparison code with nullable support
func (g *Generator) writeArrayComparisonWithNullable(
	b *strings.Builder, aArray, bArray, indent string,
	isArrayOfStructs, isArrayOfNullable, hasDepth bool,
) {
	_ = "STUB: not implemented"
	return
}

// Nullable wrapper elements - compare Null flag then Value

// Check if the wrapped value is a struct (needs Equal()) or primitive (use !=)

// Value is a struct - call Equal()

// Value is a primitive - use !=

// Array elements are structs - call Equal()

// Primitive elements can use !=
