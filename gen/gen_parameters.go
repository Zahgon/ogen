package gen

import (
	"go.uber.org/zap"

	"github.com/ogen-go/ogen/gen/ir"
	"github.com/ogen-go/ogen/openapi"
)

func vetHeaderParameterName(log *zap.Logger, name string, loc position, ignore ...string) (skip bool) {
	_ = "STUB: not implemented"
	return false
}

func (g *Generator) generateParameters(ctx *genctx, opName string, params []*openapi.Parameter) (_ []*ir.Parameter, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Path parameters are required.

// Params in different locations may have the same names,
// so we need to resolve name collision in such case.

func (g *Generator) generateParameter(ctx *genctx, opName string, p *openapi.Parameter) (ret *ir.Parameter, rerr error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Use custom name from x-ogen-name extension.

// defaultParameterJSONTag returns a default JSON Go struct tag for the given parameter type.
// Currently, returns omitempty for arrays, maps and nullable Type.GenericVariant,
// and omitzero for parameters with optional Type.GenericVariant
func defaultParameterJSONTag(t *ir.Type) string { _ = "STUB: not implemented"; return "" }

func isParamAllowed(t *ir.Type, root bool, visited map[*ir.Type]struct{}) error {
	_ = "STUB: not implemented"
	return nil
}

// Sum types are allowed in parameters.
// We'll try each variant in order during decoding.

func isSupportedParamStyle(param *openapi.Parameter) error { _ = "STUB: not implemented"; return nil }
