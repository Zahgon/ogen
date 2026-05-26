// Package parser contains the parser for OpenAPI Spec.
package parser

import (
	"net/url"

	"github.com/ogen-go/ogen"
	"github.com/ogen-go/ogen/jsonpointer"
	"github.com/ogen-go/ogen/jsonschema"
	"github.com/ogen-go/ogen/location"
	"github.com/ogen-go/ogen/openapi"
)

type refKey = jsonpointer.RefKey

type parser struct {
	// api spec, immutable.
	spec *ogen.Spec
	// root location of the spec, immutable.
	rootLoc location.Locator
	// parsed version of the spec, immutable.
	version openapi.Version

	// parsed operations.
	operations []*openapi.Operation
	// refs contains lazy-initialized referenced components.
	refs struct {
		requestBodies   map[refKey]*openapi.RequestBody
		responses       map[refKey]*openapi.Response
		parameters      map[refKey]*openapi.Parameter
		headers         map[refKey]*openapi.Header
		examples        map[refKey]*openapi.Example
		securitySchemes map[refKey]*ogen.SecurityScheme
		pathItems       map[refKey]pathItem
	}
	// securitySchemes contains security schemes defined in the root spec.
	securitySchemes map[string]*ogen.SecurityScheme
	// operationIDs holds operation IDs of already parsed operations.
	//
	// Spec says:
	//
	// 	The id MUST be unique among all operations described in the API.
	//
	// Used to detect duplicates.
	operationIDs map[string]location.Pointer

	external                     jsonschema.ExternalResolver
	rootURL                      *url.URL
	schemas                      map[string]resolver
	depthLimit                   int
	authenticationSchemes        []string
	disallowDuplicateMethodPaths bool
	rootFile                     location.File // optional, used for error messages

	schemaParser *jsonschema.Parser
}

// Parse parses raw Spec into
func Parse(spec *ogen.Spec, s Settings) (_ *openapi.API, rerr error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// pathEntry tracks a path's location and the HTTP methods defined on it.
type pathEntry struct {
	ptr     location.Pointer
	methods map[string]struct{}
}

// getPathMethods returns the set of HTTP methods defined on a PathItem.
func getPathMethods(item *ogen.PathItem) map[string]struct{} { _ = "STUB: not implemented"; return nil }

// methodsOverlap checks if two method sets have any common methods.
func methodsOverlap(a, b map[string]struct{}) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func (p *parser) parsePathItems() error {
	_ = "STUB: not implemented"

	// paths contains simple paths, e.g. "/users/{id}" -> "/users/{}".
	//
	// OpenAPI spec says:
	//
	//	Templated paths with the same hierarchy but different templated
	//	names MUST NOT exist as they are identical.
	//
	// However, when DisallowDuplicateMethodPaths is false (default),
	// we allow duplicate paths if they have different HTTP methods.
	return nil
}

// Always error if strict mode is enabled

// Check if methods overlap - if so, it's a true conflict

// Methods don't overlap - merge the method sets
