package parser

import (
	"github.com/ogen-go/ogen"
	"github.com/ogen-go/ogen/jsonschema"
	"github.com/ogen-go/ogen/location"
	"github.com/ogen-go/ogen/openapi"
)

// Expand generates an expanded ogen.Spec from given api.
func Expand(api *openapi.API) (*ogen.Spec, error) { _ = "STUB: not implemented"; return nil, nil }

type localToRemote struct {
	ref jsonschema.Ref
	ptr location.Pointer
}

type expander struct {
	components    *ogen.Components
	localToRemote map[string]localToRemote
}

func (e *expander) Spec(api *openapi.API) (spec *ogen.Spec, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *expander) Server(s openapi.Server) (expanded ogen.Server, err error) {
	_ = "STUB: not implemented"
	return *new(ogen.Server), nil
}

func (e *expander) SecurityRequirements(reqs openapi.SecurityRequirements) (expanded ogen.SecurityRequirements, err error) {
	_ = "STUB: not implemented"
	return *new(ogen.SecurityRequirements), nil
}

func (e *expander) SecurityRequirement(req openapi.SecurityRequirement) (expanded ogen.SecurityRequirement, err error) {
	_ = "STUB: not implemented"
	return *new(ogen.SecurityRequirement), nil
}

func (e *expander) SecurityScheme(s openapi.Security) (expanded *ogen.SecurityScheme, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *expander) OAuthFlows(flows *openapi.OAuthFlows) (expanded *ogen.OAuthFlows, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *expander) OAuthFlow(flow *openapi.OAuthFlow) (expanded *ogen.OAuthFlow, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *expander) Operation(op *openapi.Operation) (expanded *ogen.Operation, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *expander) Parameters(params []*openapi.Parameter) (expanded []*ogen.Parameter, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *expander) Parameter(param *openapi.Parameter) (expanded *ogen.Parameter, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *expander) genericParameter(
	param *openapi.Parameter,
	typ string,
	m map[string]*ogen.Parameter,
) (expanded *ogen.Parameter, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *expander) ParameterContent(content *openapi.ParameterContent) (map[string]ogen.Media, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *expander) RequestBody(body *openapi.RequestBody) (expanded *ogen.RequestBody, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *expander) Responses(responses openapi.Responses) (expanded ogen.Responses, err error) {
	_ = "STUB: not implemented"
	return *new(ogen.Responses), nil
}

func (e *expander) Response(resp *openapi.Response) (expanded *ogen.Response, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *expander) Headers(headers map[string]*openapi.Header) (expanded map[string]*ogen.Header, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *expander) Header(h *openapi.Header) (expanded *ogen.Header, err error) {
	_ = "STUB: not implemented"
	// Make a Paramater without "name" and "in".
	return nil, nil
}

func (e *expander) Content(content map[string]*openapi.MediaType) (expanded map[string]ogen.Media, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *expander) Media(media *openapi.MediaType) (expanded ogen.Media, err error) {
	_ = "STUB: not implemented"
	return *new(ogen.Media), nil
}

func (e *expander) Encoding(media *openapi.Encoding) (expanded ogen.Encoding, err error) {
	_ = "STUB: not implemented"
	return *new(ogen.Encoding), nil
}

func (e *expander) Schema(schema *jsonschema.Schema, walked map[*jsonschema.Schema]*ogen.Schema) (expanded *ogen.Schema, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *expander) Schemas(schemas []*jsonschema.Schema, walked map[*jsonschema.Schema]*ogen.Schema) (expanded []*ogen.Schema, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *expander) Discriminator(d *jsonschema.Discriminator) (expanded *ogen.Discriminator, _ error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *expander) XML(xml *jsonschema.XML) (expanded *ogen.XML, _ error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *expander) generateComponentName(ref jsonschema.Ref) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (e *expander) generateComponentLocalRef(
	prefix string,
	ref jsonschema.Ref,
	parentPtr location.Pointer,
) (localRef, name string, err error) {
	_ = "STUB: not implemented"
	return "", "", nil
}
