package jsonschema

import (
	"github.com/go-faster/yaml"
)

// RawProperty is item of RawProperties.
type RawProperty struct {
	Name   string
	Schema *RawSchema
}

// RawProperties is unparsed JSON Schema properties validator description.
type RawProperties []RawProperty

// MarshalYAML implements yaml.Marshaler.
func (p RawProperties) MarshalYAML() (any, error) { _ = "STUB: not implemented"; return *new(any), nil }

// UnmarshalYAML implements yaml.Unmarshaler.
func (p *RawProperties) UnmarshalYAML(node *yaml.Node) error { _ = "STUB: not implemented"; return nil }

// MarshalJSON implements json.Marshaler.
func (p RawProperties) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// UnmarshalJSON implements json.Unmarshaler.
func (p *RawProperties) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

// AdditionalProperties is JSON Schema additionalProperties validator description.
type AdditionalProperties struct {
	Bool   *bool
	Schema RawSchema
}

// MarshalYAML implements yaml.Marshaler.
func (p AdditionalProperties) MarshalYAML() (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

// UnmarshalYAML implements yaml.Unmarshaler.
func (p *AdditionalProperties) UnmarshalYAML(node *yaml.Node) error {
	_ = "STUB: not implemented"
	return nil
}

// MarshalJSON implements json.Marshaler.
func (p AdditionalProperties) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (p *AdditionalProperties) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// RawPatternProperty is item of RawPatternProperties.
type RawPatternProperty struct {
	Pattern string
	Schema  *RawSchema
}

// RawPatternProperties is unparsed JSON Schema patternProperties validator description.
type RawPatternProperties []RawPatternProperty

// MarshalYAML implements yaml.Marshaler.
func (p RawPatternProperties) MarshalYAML() (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

// UnmarshalYAML implements yaml.Unmarshaler.
func (p *RawPatternProperties) UnmarshalYAML(node *yaml.Node) error {
	_ = "STUB: not implemented"
	return nil
}

// MarshalJSON implements json.Marshaler.
func (p RawPatternProperties) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (p *RawPatternProperties) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// RawItems is unparsed JSON Schema items validator description.
type RawItems struct {
	Item  *RawSchema
	Items []*RawSchema
}

// MarshalYAML implements yaml.Marshaler.
func (p RawItems) MarshalYAML() (any, error) { _ = "STUB: not implemented"; return *new(any), nil }

// UnmarshalYAML implements yaml.Unmarshaler.
func (p *RawItems) UnmarshalYAML(node *yaml.Node) error { _ = "STUB: not implemented"; return nil }

// MarshalJSON implements json.Marshaler.
func (p RawItems) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// UnmarshalJSON implements json.Unmarshaler.
func (p *RawItems) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }
