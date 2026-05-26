package jsonschema

import (
	"github.com/go-faster/yaml"
)

// RootResolver is ReferenceResolver implementation.
type RootResolver struct {
	root *yaml.Node
}

// NewRootResolver creates new RootResolver.
func NewRootResolver(root *yaml.Node) *RootResolver { _ = "STUB: not implemented"; return nil }

// ResolveReference implements ReferenceResolver.
func (r *RootResolver) ResolveReference(ref string) (rawSchema *RawSchema, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}
