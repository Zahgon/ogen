package parser

import (
	"github.com/go-faster/yaml"

	"github.com/ogen-go/ogen"
	"github.com/ogen-go/ogen/jsonpointer"
	"github.com/ogen-go/ogen/location"
	"github.com/ogen-go/ogen/openapi"
)

type resolver struct {
	node *yaml.Node
	file location.File
}

func (p *parser) getResolver(loc string) (r resolver, rerr error) {
	_ = "STUB: not implemented"
	return *new(resolver), nil
}

func resolvePointer(root *yaml.Node, ptr string, to any) error {
	_ = "STUB: not implemented"
	return nil
}

// componentResolve contains all the information needed to resolve a component.
type componentResolve[Raw, Target any] struct {
	// prefix is the usual prefix of the component reference (e.g "#/components/requestBodies/").
	prefix string
	// components is the root spec components field.
	components map[string]Raw
	// cache is the cache of already resolved components.
	cache map[refKey]Target
	// parse is the function that parses the raw component.
	parse func(Raw, *jsonpointer.ResolveCtx) (Target, error)
}

// resolveComponent is a generic function that resolves a component.
//
// We return a boolean indicating whether the component was already cached.
// Do not set the ref if it was cached.
func resolveComponent[Raw, Target any](
	p *parser,
	cr componentResolve[Raw, Target],
	ref string,
	ctx *jsonpointer.ResolveCtx,
) (key refKey, zero Target, cached bool, _ error) {
	_ = "STUB: not implemented"
	return *new(refKey), *new(Target), false, nil
}

func (p *parser) resolveCtx() *jsonpointer.ResolveCtx { _ = "STUB: not implemented"; return nil }

func (p *parser) resolveRequestBody(ref string, ctx *jsonpointer.ResolveCtx) (*openapi.RequestBody, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *parser) resolveResponse(ref string, ctx *jsonpointer.ResolveCtx) (*openapi.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *parser) resolveParameter(ref string, ctx *jsonpointer.ResolveCtx) (*openapi.Parameter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *parser) resolveHeader(headerName, ref string, ctx *jsonpointer.ResolveCtx) (*openapi.Header, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *parser) resolveExample(ref string, ctx *jsonpointer.ResolveCtx) (*openapi.Example, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *parser) resolveSecurityScheme(ref string, ctx *jsonpointer.ResolveCtx) (*ogen.SecurityScheme, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *parser) resolvePathItem(
	itemPath unparsedPath,
	ref string,
	ctx *jsonpointer.ResolveCtx,
) (pathItem, error) {
	_ = "STUB: not implemented"
	return *new(pathItem), nil
}
