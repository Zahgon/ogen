package parser

import (
	"github.com/go-faster/errors"

	"github.com/ogen-go/ogen/openapi"
)

type pathParser[P any] struct {
	// Input path.
	path string // immutable
	// Callback to lookup parameter by name.
	lookup func(name string) (P, bool) // immutable

	// Parser state.
	parts []openapi.PathPart[P] // parsed parts
	part  []rune                // current part
	param bool                  // current part is param name?
}

func pathID(path string) (string, error) { _ = "STUB: not implemented"; return "", nil }

var errInvalidPathUTF8 = errors.New("path must be valid UTF-8 string")

func parsePath(path string, params []*openapi.Parameter) (openapi.Path, error) {
	_ = "STUB: not implemented"
	return *new(openapi.Path), nil
}

// Validate and unescape path.
//
// FIXME(tdakkota): OpenAPI spec, as always, is not clear about path validation.
//  All we know is that it MUST start with a slash.
// 	At the same time, https://swagger.io/docs/specification/paths-and-operations/ says that
// 	paths must not include query parameters.
//  In summary, we do not pass URL scheme, user info, host or query string.
//

func parseServerURL(u string, lookup func(name string) (openapi.ServerVariable, bool)) (openapi.ServerURL, error) {
	_ = "STUB: not implemented"
	return *new(openapi.ServerURL), nil
}

func (p *pathParser[P]) Parse() ([]openapi.PathPart[P], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *pathParser[P]) parse() error { _ = "STUB: not implemented"; return nil }

type pathParameterNotSpecifiedError struct {
	Name string
}

func (p *pathParameterNotSpecifiedError) Error() string { _ = "STUB: not implemented"; return "" }

func (p *pathParser[P]) push() error { _ = "STUB: not implemented"; return nil }
