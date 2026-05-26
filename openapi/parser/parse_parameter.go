package parser

import (
	"github.com/ogen-go/ogen"
	"github.com/ogen-go/ogen/jsonpointer"
	"github.com/ogen-go/ogen/location"
	"github.com/ogen-go/ogen/openapi"
)

func canonicalParamName(name string, in openapi.ParameterLocation) string {
	_ = "STUB: not implemented"
	return ""
}

func mergeParams(opParams, itemParams []*openapi.Parameter) []*openapi.Parameter {
	_ = "STUB: not implemented"
	return nil
}

// Param defined in operation take precedence over param defined in pathItem.

func (p *parser) parseParams(
	params []*ogen.Parameter,
	locator location.Locator,
	ctx *jsonpointer.ResolveCtx,
) ([]*openapi.Parameter, error) {
	_ = "STUB: not implemented"
	// Unique parameter is defined by a combination of a name and location.
	return nil, nil
}

func (p *parser) validateParameter(
	name string,
	locatedIn openapi.ParameterLocation,
	param *ogen.Parameter,
	file location.File,
) error {
	_ = "STUB: not implemented"
	return nil
}

// https://github.com/OAI/OpenAPI-Specification/discussions/2875

// Path parameters are always required.

func (p *parser) parseParameter(param *ogen.Parameter, ctx *jsonpointer.ResolveCtx) (_ *openapi.Parameter, rerr error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Parse x-ogen-name extension.

func inferParamStyle(locatedIn openapi.ParameterLocation, style string) openapi.ParameterStyle {
	_ = "STUB: not implemented"
	return *new(openapi.ParameterStyle)
}

func inferParamExplode(locatedIn openapi.ParameterLocation, explode *bool) bool {
	_ = "STUB: not implemented"
	return false
}

// When style is form, the default value is true.
// For all other styles, the default value is false.

func (p *parser) validateParamStyle(param *openapi.Parameter, file location.File) error {
	_ = "STUB: not implemented"
	// https://swagger.io/docs/specification/serialization/
	return nil
}
