package jsonschema

import (
	"encoding/json"

	"github.com/go-faster/yaml"
)

type (
	// RawValue is a raw JSON value.
	RawValue json.RawMessage
	// Default is a default value.
	Default = RawValue
	// Example is an example value.
	Example = RawValue
)

// MarshalYAML implements yaml.Marshaler.
func (n RawValue) MarshalYAML() (any, error) { _ = "STUB: not implemented"; return *new(any), nil }

// UnmarshalYAML implements yaml.Unmarshaler.
func (n *RawValue) UnmarshalYAML(node *yaml.Node) error { _ = "STUB: not implemented"; return nil }

// MarshalJSON implements json.Marshaler.
func (n RawValue) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// UnmarshalJSON implements json.Unmarshaler.
func (n *RawValue) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

func convertJSONToRawYAML(raw json.RawMessage) (_ *yaml.Node, rerr error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func convertYAMLtoRawJSON(node *yaml.Node) (_ json.RawMessage, rerr error) {
	_ = "STUB: not implemented"
	return *new(json.RawMessage), nil
}
