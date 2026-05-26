package jsonschema

import (
	"encoding/json"

	"github.com/go-faster/yaml"
)

// Num represents JSON number.
type Num json.RawMessage

// MarshalYAML implements yaml.Marshaler.
func (n Num) MarshalYAML() (any, error) { _ = "STUB: not implemented"; return *new(any), nil }

// UnmarshalYAML implements yaml.Unmarshaler.
func (n *Num) UnmarshalYAML(node *yaml.Node) error { _ = "STUB: not implemented"; return nil }

// MarshalJSON implements json.Marshaler.
func (n Num) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// UnmarshalJSON implements json.Unmarshaler.
func (n *Num) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }
