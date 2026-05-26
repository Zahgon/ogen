package gen

import (
	"github.com/ogen-go/ogen/gen/ir"
	"github.com/ogen-go/ogen/openapi"
)

func (g *Generator) generateSecurityAPIKey(
	s *ir.Security,
	operationName string,
	spec openapi.SecurityScheme,
) (*ir.Security, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (g *Generator) generateSecurityOauth2(
	s *ir.Security,
	operationName string,
	spec openapi.SecurityScheme,
) *ir.Security {
	_ = "STUB: not implemented"
	return nil
}

func (g *Generator) generateSecurityHTTP(
	s *ir.Security,
	operationName string,
	spec openapi.SecurityScheme,
) (*ir.Security, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (g *Generator) generateCustomSecurity(
	s *ir.Security,
	operationName string,
	spec openapi.SecurityScheme,
) *ir.Security {
	_ = "STUB: not implemented"
	return nil
}

func (g *Generator) generateSecurity(ctx *genctx, operationName string, spec openapi.SecurityScheme) (r *ir.Security, rErr error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (g *Generator) generateSecurities(
	ctx *genctx,
	operationName string,
	spec openapi.SecurityRequirements,
) (r ir.SecurityRequirements, _ error) {
	_ = "STUB: not implemented"
	return *new(ir.SecurityRequirements), nil
}

// Skip entire requirement if at least one security is not implemented.
