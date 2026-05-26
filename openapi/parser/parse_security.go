package parser

import (
	"github.com/ogen-go/ogen"
	"github.com/ogen-go/ogen/jsonpointer"
	"github.com/ogen-go/ogen/location"
	"github.com/ogen-go/ogen/openapi"
)

func (p *parser) parseSecurityScheme(
	scheme *ogen.SecurityScheme,
	ctx *jsonpointer.ResolveCtx,
) (_ *ogen.SecurityScheme, rerr error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func forEachFlow(flows *ogen.OAuthFlows, cb func(flow *ogen.OAuthFlow, authURL, tokenURL bool) error) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *parser) validateOAuthFlows(flows *ogen.OAuthFlows, file location.File) (rerr error) {
	_ = "STUB: not implemented"
	return nil
}

func cloneOAuthFlows(flows ogen.OAuthFlows, file location.File) (r openapi.OAuthFlows) {
	_ = "STUB: not implemented"
	return *new(openapi.OAuthFlows)
}

func (p *parser) parseSecurityRequirementScheme(name string, scheme *ogen.SecurityScheme) (openapi.Security, error) {
	_ = "STUB: not implemented"
	// Note that we use root context/file.
	return *new(openapi.Security), nil
}

func (p *parser) parseSecurityRequirements(
	requirements ogen.SecurityRequirements,
	locator location.Locator,
	ctx *jsonpointer.ResolveCtx,
) (openapi.SecurityRequirements, error) {
	_ = "STUB: not implemented"
	return *new(openapi.SecurityRequirements), nil
}
