package ir

import (
	"github.com/ogen-go/ogen/jsonschema"
)

// InlineField defines how to inline field.
type InlineField int

const (
	InlineNone InlineField = iota
	InlineAdditional
	InlinePattern
	InlineSum
)

// Field of structure.
type Field struct {
	// Go Name of field.
	Name string
	// Type of field.
	Type *Type
	// JSON tag. May be empty.
	Tag Tag
	// Whether field is inlined map (i.e. additionalProperties, patternProperties).
	Inline InlineField
	// Spec is property schema. May be nil.
	Spec *jsonschema.Property
}

// ValidationName returns name for FieldError.
func (f Field) ValidationName() string { _ = "STUB: not implemented"; return "" }

// Default returns default value of this field, if it is set.
func (f Field) Default() Default { _ = "STUB: not implemented"; return *new(Default) }

// Const returns const value of this field, if it is set.
func (f Field) Const() Const { _ = "STUB: not implemented"; return *new(Const) }

// GoDoc returns field godoc.
func (f Field) GoDoc() []string { _ = "STUB: not implemented"; return nil }

// DefaultFields returns fields with default values.
func (t Type) DefaultFields() (r []*Field) { _ = "STUB: not implemented"; return nil }

// HasDefaultFields whether type has fields with default values.
func (t Type) HasDefaultFields() bool { _ = "STUB: not implemented"; return false }

func (t Type) parameters(keep func(t *Type) bool) (params []Parameter) {
	_ = "STUB: not implemented"
	return nil
}

func (t Type) FormParameters() (params []Parameter) { _ = "STUB: not implemented"; return nil }

func (t Type) FileParameters() (params []Parameter) { _ = "STUB: not implemented"; return nil }
