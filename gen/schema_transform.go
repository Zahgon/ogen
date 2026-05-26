package gen

import (
	"github.com/ogen-go/ogen/jsonschema"
)

func transformSchema(schema *jsonschema.Schema) *jsonschema.Schema {
	_ = "STUB: not implemented"
	return nil
}

// transformSingleOneOf detects and handles single oneOf patterns.
//
// single oneOf pattern is:
//
//	oneOf:
//	  - <schema>
//
// if such pattern is detected, this function will return the inner schema.
func transformSingleOneOf(schema *jsonschema.Schema) *jsonschema.Schema {
	_ = "STUB: not implemented"
	return nil
}

// transformNullableUnionType detects and handles nullable oneOf/anyOf patterns.
//
// nullable oneOf pattern is:
//
//	oneOf:
//	  - type: "null"
//	  - <schema>
//
// or
//
//	anyOf:
//	  - type: "null"
//	  - <schema>
//
// if such pattern is detected, this function will return a Nulllable version of the inner schema.
func transformNullableUnionType(schema *jsonschema.Schema) *jsonschema.Schema {
	_ = "STUB: not implemented"
	return nil
}

// If we didn't find exactly one null and one non-null variant, don't handle

// Return nullable version of the underlined schema.
// Make a shallow copy to avoid mutating the original schema.
