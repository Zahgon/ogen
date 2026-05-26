package gen

import (
	"go.uber.org/zap"

	"github.com/ogen-go/ogen/gen/ir"
	"github.com/ogen-go/ogen/openapi"
)

func (g *Generator) generateOperation(ctx *genctx, webhookName string, spec *openapi.Operation) (_ *ir.Operation, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Convert []openapi.Parameter to []*ir.Parameter.

// Convert []openapi.PathPart to []*ir.PathPart

func vetPathParametersUsed(log *zap.Logger, parts openapi.Path, params []*openapi.Parameter) {
	_ = "STUB: not implemented"
	return
}

func convertPathParts(parts openapi.Path, params []*ir.Parameter) []*ir.PathPart {
	_ = "STUB: not implemented"
	return nil
}
