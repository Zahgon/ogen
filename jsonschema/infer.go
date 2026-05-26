package jsonschema

import (
	"github.com/go-faster/jx"
)

// Infer returns a JSON Schema that is inferred from the given JSON.
type Infer struct {
	target RawSchema
}

// Target returns the target schema.
func (i Infer) Target() RawSchema {
	_ = "STUB: not implemented"

	// Apply applies given data to the schema state.
	return *new(RawSchema)
}

func (i *Infer) Apply(data []byte) error { _ = "STUB: not implemented"; return nil }

func applyType(s *RawSchema, tt string) { _ = "STUB: not implemented"; return }

func hasType(s *RawSchema, tt string) bool { _ = "STUB: not implemented"; return false }

func replaceType(s *RawSchema, from, to string) bool { _ = "STUB: not implemented"; return false }

func apply(s *RawSchema, d *jx.Decoder) error { _ = "STUB: not implemented"; return nil }

// Set s.Properties to non-nil slice to mark that it is not first apply.

// Collect existing properties.

// Collect required properties.

// If it is the first apply, mark property as required.

// Delete required properties that are not in this object.

// Write required properties.

// Sort fields to make output deterministic.
