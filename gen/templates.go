package gen

import (
	"embed"
	"sync"
	"text/template"

	"github.com/ogen-go/ogen/gen/ir"
)

// OperationElem is variable name for generating per-operation functions.
type OperationElem struct {
	// Operation is the operation.
	Operation *ir.Operation
	// Config is the template configuration.
	Config TemplateConfig
}

// RouterElem is variable helper for router generation.
type RouterElem struct {
	// ParameterIndex is index of parameter of this route part.
	ParameterIndex int
	Route          *RouteNode
}

// DefaultElem is variable helper for setting default values.
type DefaultElem struct {
	// Type is type of this DefaultElem.
	Type *ir.Type
	// Var is decoding/encoding variable Go name (obj) or selector (obj.Field).
	Var string
	// Default is default value to set.
	Default ir.Default
}

// Elem is variable helper for recursive array or object encoding or decoding.
type Elem struct {
	// Sub whether this Elem has parent Elem.
	Sub bool
	// Type is type of this Elem.
	Type *ir.Type
	// Var is decoding/encoding variable Go name (obj) or selector (obj.Field).
	Var string
	// Tag contains info about field tags, if any.
	Tag ir.Tag
	// First whether this field is first.
	First bool
}

// NextVar returns name of variable for decoding recursive call.
//
// Needed to make variable names unique.
func (e Elem) NextVar() string {
	_ = "STUB: not implemented"

	// No recursion, returning default name.
	return ""
}

type ResponseElem struct {
	Response *ir.Response
	Ptr      bool
}

// templateFunctions returns functions which used in templates.
func templateFunctions() template.FuncMap { _ = "STUB: not implemented"; return *new(template.FuncMap) }

// Helpers for recursive encoding and decoding.

// Recursive array element (e.g. array of arrays).

// Initial array element.

// Field of structure.

// We use any to prevent template type matching errors
// for type aliases (e.g. for quoting ir.ContentType).

// Fast path for string.

// Fast path for string.

//go:embed _template/*
var templates embed.FS

var _templates struct {
	sync.Once
	val *template.Template
}

// vendoredTemplates parses and returns vendored code generation templates.
func vendoredTemplates() *template.Template { _ = "STUB: not implemented"; return nil }

func isObjectParam(p *ir.Parameter) bool { _ = "STUB: not implemented"; return false }

// "content" encoding.

func paramObjectFields(typ *ir.Type) string { _ = "STUB: not implemented"; return "" }

// uniqueResponseTypes deduplicates response types by Type.Name to avoid duplicate case statements
// in type switches. When multiple responses share the same type (e.g., multiple patterns using the
// same schema), we only need one case statement.
func uniqueResponseTypes(responses []ir.ResponseInfo) []ir.ResponseInfo {
	_ = "STUB: not implemented"
	return nil
}

// Raw responses are handled separately in the template

// dedupeVariantsByType deduplicates variants by their FieldType to avoid duplicate type checks.
// When multiple variants have the same field type (or no type discrimination), keep only unique entries.
func dedupeVariantsByType(variants []ir.UniqueFieldVariant) []ir.UniqueFieldVariant {
	_ = "STUB: not implemented"
	return nil
}

// If FieldType is empty (no type discrimination), include all variants

// needsArrayElementDiscrimination checks if all variants have the same jx.Array FieldType
// but different ArrayElementTypes, requiring element-level discrimination.
func needsArrayElementDiscrimination(variants []ir.UniqueFieldVariant) bool {
	_ = "STUB: not implemented"
	return false
}

// All variants must be arrays

// Count unique element types

// dedupeVariantsByArrayElementType deduplicates array variants by their ArrayElementType.
// Used when all variants are arrays that need element-level discrimination.
func dedupeVariantsByArrayElementType(variants []ir.UniqueFieldVariant) []ir.UniqueFieldVariant {
	_ = "STUB: not implemented"
	return nil
}

// If ArrayElementType is empty, include the variant
