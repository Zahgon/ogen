package jsonschema

import (
	"encoding/json"
)

func inferJSONType(v json.RawMessage) (string, error) { _ = "STUB: not implemented"; return "", nil }

func parseEnumValues(s *Schema, rawValues []json.RawMessage) ([]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseJSONValue(root *Schema, v json.RawMessage) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

// See https://github.com/OAI/OpenAPI-Specification/blob/main/proposals/2019-10-31-Clarify-Nullable.md#if-a-schema-specifies-nullable-true-and-enum-1-2-3-does-that-schema-allow-null-values-see-1900.
func handleNullableEnum(s *Schema) {
	_ = "STUB: not implemented"
	// Workaround: handle nullable enums correctly.
	//
	// Notice that nullable enum requires `null` in value list.
	//
	// Check that enum contains `null` value.
	return
}

// Filter all `null`s.
