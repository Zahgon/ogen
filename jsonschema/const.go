package jsonschema

import (
	"encoding/json"

	"github.com/go-faster/yaml"
)

// Const is JSON Schema const validator description.
type Const json.RawMessage

// MarshalYAML implements yaml.Marshaler.
func (c Const) MarshalYAML() (any, error) { _ = "STUB: not implemented"; return *new(any), nil }

// UnmarshalYAML implements yaml.Unmarshaler.
func (c *Const) UnmarshalYAML(node *yaml.Node) error { _ = "STUB: not implemented"; return nil }

// MarshalJSON implements json.Marshaler.
func (c Const) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// UnmarshalJSON implements json.Unmarshaler.
func (c *Const) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

// validateConst rejects empty objects.
func validateConst(raw []byte) error { _ = "STUB: not implemented"; return nil }
