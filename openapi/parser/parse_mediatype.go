package parser

import (
	"github.com/ogen-go/ogen"
	"github.com/ogen-go/ogen/jsonpointer"
	"github.com/ogen-go/ogen/location"
	"github.com/ogen-go/ogen/openapi"
)

func (p *parser) parseContent(
	content map[string]ogen.Media,
	locator location.Locator,
	ctx *jsonpointer.ResolveCtx,
) (_ map[string]*openapi.MediaType, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *parser) parseParameterContent(
	content map[string]ogen.Media,
	locator location.Locator,
	ctx *jsonpointer.ResolveCtx,
) (*openapi.ParameterContent, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Report only 2 entries.

// Set message for the first position.

func (p *parser) parseMediaType(ct string, m ogen.Media, ctx *jsonpointer.ResolveCtx) (_ *openapi.MediaType, rerr error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// OpenAPI 3.0.3 doc says:
//
//   Furthermore, referencing a schema which contains an example,
//   the example value SHALL override the example provided by the schema.
//
// Probably this will be rewritten later.
// Kept for backward compatibility.
