package parser

import (
	"github.com/ogen-go/ogen"
	"github.com/ogen-go/ogen/jsonschema"
)

type componentsResolver struct {
	components map[string]*ogen.Schema
	root       *jsonschema.RootResolver
}

func (c componentsResolver) ResolveReference(ref string) (*jsonschema.RawSchema, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
