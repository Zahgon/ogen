package jsonschema

import (
	"github.com/ogen-go/ogen/jsonpointer"
	"github.com/ogen-go/ogen/location"
)

// ReferenceResolver resolves JSON schema references.
type ReferenceResolver interface {
	ResolveReference(ref string) (*RawSchema, error)
}

type resolver struct {
	ReferenceResolver
	file location.File
}

func (p *Parser) getResolver(loc string) (r resolver, rerr error) {
	_ = "STUB: not implemented"
	return *new(resolver), nil
}

func (p *Parser) resolve(ref string, ctx *jsonpointer.ResolveCtx) (_ *Schema, rerr error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Drop the resolved ref to prevent false-positive infinite recursion detection.
