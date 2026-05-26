// Package jsonschema contains parser for JSON Schema.
package jsonschema

import (
	"github.com/ogen-go/ogen/jsonpointer"
	"github.com/ogen-go/ogen/location"
)

const (
	xOgenName       = "x-ogen-name"
	xOgenProperties = "x-ogen-properties"
	xOgenType       = "x-ogen-type"
	xOgenTimeFormat = "x-ogen-time-format"
	xOapiExtraTags  = "x-oapi-codegen-extra-tags"
	xOgenValidate   = "x-ogen-validate"
)

// Parser parses JSON schemas.
type Parser struct {
	external ExternalResolver
	schemas  map[string]resolver
	refcache map[jsonpointer.RefKey]*Schema

	rootFile location.File // optional, used for error messages

	inferTypes                bool
	allowCrossTypeConstraints bool
}

// NewParser creates new Parser.
func NewParser(s Settings) *Parser { _ = "STUB: not implemented"; return nil }

// Parse parses given RawSchema and returns parsed Schema.
func (p *Parser) Parse(schema *RawSchema, ctx *jsonpointer.ResolveCtx) (*Schema, error) {
	_ = "STUB: not implemented"
	return nil,

		// Resolve resolves Schema by given ref.
		nil
}

func (p *Parser) Resolve(ref string, ctx *jsonpointer.ResolveCtx) (*Schema, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *Parser) parse(schema *RawSchema, ctx *jsonpointer.ResolveCtx) (_ *Schema, rerr error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *Parser) parse1(schema *RawSchema, ctx *jsonpointer.ResolveCtx, hook func(*Schema) *Schema) (*Schema, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// FIXME(tdakkota): linear search

// validateGoIdent checks that given ident is valid and is exported.
func validateGoIdent(ident string) error {
	_ = "STUB: not implemented"
	// TODO(tdakkota): move to generator package?
	//
	//	For now, keep as part of parser to use user-friendly location errors
	return nil
}

func (p *Parser) parseSchema(schema *RawSchema, ctx *jsonpointer.ResolveCtx, hook func(*Schema) *Schema) (_ *Schema, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Try to infer schema type from properties.

// FIXME(tdakkota): check for existence instead of true?

// Generic fields.

// Allow extra fields.

// Allow cross-type constraints.
// They will be interpreted during validation code generation.

// Object validators

// Array validators

// Number validators

// String validators

// Object properties

// Array properties

// Integer, Number properties

// The value of "multipleOf" MUST be a number, strictly greater than 0.

// String properties

func (p *Parser) parseMany(schemas []*RawSchema, loc location.Locator, ctx *jsonpointer.ResolveCtx) ([]*Schema, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *Parser) extendInfo(schema *RawSchema, s *Schema, file location.File) *Schema {
	_ = "STUB: not implemented"
	return nil
}

// Nullable enums will be handled later.

func (p *Parser) parseDiscriminator(d *RawDiscriminator, ctx *jsonpointer.ResolveCtx) (_ *Discriminator, rerr error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// See https://github.com/OAI/OpenAPI-Specification/issues/2520#issuecomment-1139961158.

// JSON Reference usually contains a fragment, e.g. "#/components/schemas/Foo" or
// "foo.json#/definitions/Bar", but this looks like a schema name.
//
// Try to find it in the components, if it is root spec.

// It seems there is no schema with such name, try to resolve as a plain JSON Reference.
