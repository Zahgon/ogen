package parser

import (
	"github.com/ogen-go/ogen"
	"github.com/ogen-go/ogen/jsonpointer"
	"github.com/ogen-go/ogen/openapi"
)

func (p *parser) parseRequestBody(body *ogen.RequestBody, ctx *jsonpointer.ResolveCtx) (_ *openapi.RequestBody, rerr error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// See https://github.com/OAI/OpenAPI-Specification/discussions/2875.
