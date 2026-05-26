package parser

import (
	"github.com/ogen-go/ogen"
	"github.com/ogen-go/ogen/jsonpointer"
	"github.com/ogen-go/ogen/jsonschema"
	"github.com/ogen-go/ogen/location"
	"github.com/ogen-go/ogen/openapi"
)

type (
	pathItem     = []*openapi.Operation
	unparsedPath struct {
		path string
		loc  location.Locator
		file location.File
	}
)

const (
	xOgenOperationGroup = "x-ogen-operation-group"
)

func (up unparsedPath) String() string { _ = "STUB: not implemented"; return "" }

func (p *parser) parsePathItem(
	up unparsedPath,
	item *ogen.PathItem,
	ctx *jsonpointer.ResolveCtx,
) (_ pathItem, rerr error) {
	_ = "STUB: not implemented"
	return *new(pathItem), nil
}

// Look for x-ogen-operation-group on the PathItem.
// Use it as a default value for operations.

// Validate that additionalOperations don't contain any
// any entry for the methods that can be defined by other fixed fields.

func (p *parser) parseOp(
	up unparsedPath,
	httpMethod string,
	spec ogen.Operation,
	itemParams []*openapi.Parameter,
	ctx *jsonpointer.ResolveCtx,
	operationGroup string,
) (_ *openapi.Operation, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Merge operation parameters with pathItem parameters.

// Special case: point to the operation "parameters" what caused the error.
// It is helpful, since one Path Item may contain multiple operations.

// Use operation level security.

func forEachOps(item *ogen.PathItem, f func(method string, op ogen.Operation) error) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *parser) parseOperationGroup(common jsonschema.OpenAPICommon, operationGroup *string) error {
	_ = "STUB: not implemented"
	return nil
}
