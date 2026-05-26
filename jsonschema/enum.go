package jsonschema

import (
	"encoding/json"

	"github.com/go-faster/yaml"
)

// Enum is JSON Schema enum validator description.
type Enum []json.RawMessage

// MarshalYAML implements yaml.Marshaler.
func (n Enum) MarshalYAML() (any, error) { _ = "STUB: not implemented"; return *new(any), nil }

// UnmarshalYAML implements yaml.Unmarshaler.
func (n *Enum) UnmarshalYAML(node *yaml.Node) error { _ = "STUB: not implemented"; return nil }
