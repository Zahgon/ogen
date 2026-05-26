package parser

import (
	"regexp"

	"github.com/ogen-go/ogen"
	"github.com/ogen-go/ogen/location"
	"github.com/ogen-go/ogen/openapi"
)

var componentsKeyRegex = regexp.MustCompile(`^[a-zA-Z0-9.\-_]+$`)

// validateComponentsKey validates components key.
//
// Spec says:
//
//	All the fixed fields declared above are objects that MUST use keys that
//	match the regular expression: ^[a-zA-Z0-9\.\-_]+$.
//
// See https://spec.openapis.org/oas/v3.1.0#components-object.
func validateComponentsKey[Object any](p *parser, m map[string]Object, locator location.Locator) error {
	_ = "STUB: not implemented"
	return nil
}

// validateComponentsKeys validates components keys.
//
// See validateComponentsKey comment.
func validateComponentsKeys(p *parser, c *ogen.Components) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *parser) parseComponents(c *ogen.Components) (_ *openapi.Components, rerr error) {
	_ = "STUB: not implemented"
	return nil, nil
}
