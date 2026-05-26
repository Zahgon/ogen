package gen

import (
	"github.com/ogen-go/ogen/gen/ir"
	"github.com/ogen-go/ogen/jsonschema"
	"github.com/ogen-go/ogen/openapi"
)

func (g *Generator) generateResponses(ctx *genctx, opName string, responses openapi.Responses) (_ *ir.Responses, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Sort responses by status code.

// We'll need to use an interface-based approach for raw responses

// Add raw response concrete types for content types with RawResponse=true

// addRawResponseTypes adds concrete types for raw responses that implement the interface
func addRawResponseTypes(ctx *genctx, result *ir.Responses, iface *ir.Type, opName string) error {
	_ = "STUB: not implemented"
	return nil
}

// Remove the original structured type from the interface
// since we're replacing it with a raw response type

func (g *Generator) responseToIR(
	ctx *genctx,
	name, doc string,
	resp *openapi.Response,
	withStatusCode bool,
) (ret *ir.Response, rerr error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Check for unsupported response content types.

// Use content-type-specific name for wrapper when there are multiple contents
// to avoid name conflicts (e.g., when both application/json and
// application/vnd.github.v3.star+json have array schemas without names).

func wrapResponseType(
	ctx *genctx,
	name string,
	respRef jsonschema.Ref,
	t *ir.Type,
	headers map[string]*ir.Parameter,
	withStatusCode bool,
	multipleContents bool,
) (ret *ir.Type, rerr error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Prefer response name to schema name in case of wrapping.

func injectHeaderFields(headers map[string]*ir.Parameter, t *ir.Type) {
	_ = "STUB: not implemented"
	return
}
