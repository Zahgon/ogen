package parser

import (
	"github.com/ogen-go/ogen"
	"github.com/ogen-go/ogen/jsonpointer"
	"github.com/ogen-go/ogen/location"
	"github.com/ogen-go/ogen/openapi"
)

func (p *parser) parseResponses(
	responses ogen.Responses,
	locator location.Locator,
	ctx *jsonpointer.ResolveCtx,
) (result openapi.Responses, _ error) {
	_ = "STUB: not implemented"
	return *new(openapi.Responses), nil
}

func (p *parser) parseResponse(resp *ogen.Response, ctx *jsonpointer.ResolveCtx) (_ *openapi.Response, rerr error) {
	_ = "STUB: not implemented"
	return nil, nil
}
