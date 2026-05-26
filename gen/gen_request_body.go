package gen

import (
	"github.com/ogen-go/ogen/gen/ir"
	"github.com/ogen-go/ogen/openapi"
)

func (g *Generator) generateRequest(ctx *genctx, opName string, body *openapi.RequestBody) (*ir.Request, error) {
	_ = "STUB: not implemented"
	return nil,

		// Filter early to check the number of media below.
		//
		// FIXME(tdakkota): do not modify the original body.
		nil
}

// Generate optional type only if there is only one media type and body is not required.
//
// Otherwise, we generate a special "EmptyBody" case.

// Generate an empty body case only if there is more than one media type.
//
// If there is only one media type, we generate an optional type instead.
