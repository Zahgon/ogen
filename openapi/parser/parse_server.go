package parser

import (
	"github.com/ogen-go/ogen"
	"github.com/ogen-go/ogen/jsonpointer"
	"github.com/ogen-go/ogen/openapi"
)

func (p *parser) parseServers(servers []ogen.Server, ctx *jsonpointer.ResolveCtx) ([]openapi.Server, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *parser) parseServer(
	s ogen.Server,
	dedup map[string]struct{},
	ctx *jsonpointer.ResolveCtx,
) (_ openapi.Server, rerr error) {
	_ = "STUB: not implemented"
	return *new(openapi.Server), nil
}

// Validate variables.

// Validate URL.

// Parse extensions.
//
// TODO(tdakkota): describe extensions somewhere, it would be nice to have machine-readable
// 	description of extensions, their types, and their validation rules.

// Ensure that ${name}Server is a valid Go identifier.
