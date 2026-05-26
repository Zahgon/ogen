package jsonschema

import (
	"github.com/go-faster/yaml"

	"github.com/ogen-go/ogen/location"
)

// Extensions map is "^x-" fields list.
type Extensions map[string]yaml.Node

func isExtensionKey(key string) bool { _ = "STUB: not implemented"; return false }

// MarshalYAML implements yaml.Marshaler.
func (p Extensions) MarshalYAML() (any, error) { _ = "STUB: not implemented"; return *new(any), nil }

// UnmarshalYAML implements yaml.Unmarshaler.
func (p *Extensions) UnmarshalYAML(node *yaml.Node) error { _ = "STUB: not implemented"; return nil }

// FIXME(tdakkota): use *yamlx.Node instead of yaml.Node

// MarshalJSON implements json.Marshaler.
func (p Extensions) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// UnmarshalJSON implements json.Unmarshaler.
func (p *Extensions) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

// OpenAPICommon is common fields for OpenAPI objects.
type OpenAPICommon struct {
	Extensions
	location.Locator
}

// MarshalYAML implements yaml.Marshaler.
func (p OpenAPICommon) MarshalYAML() (any, error) { _ = "STUB: not implemented"; return *new(any), nil }

// UnmarshalYAML implements yaml.Unmarshaler.
func (p *OpenAPICommon) UnmarshalYAML(node *yaml.Node) error { _ = "STUB: not implemented"; return nil }

// MarshalJSON implements json.Marshaler.
func (p OpenAPICommon) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// UnmarshalJSON implements json.Unmarshaler.
func (p *OpenAPICommon) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }
