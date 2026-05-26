package gen

import (
	"go.uber.org/zap"

	"github.com/ogen-go/ogen/gen/ir"
	"github.com/ogen-go/ogen/jsonschema"
)

func saveSchemaTypes(ctx *genctx, gen *schemaGen, refEncoding map[jsonschema.Ref]ir.Encoding) error {
	_ = "STUB: not implemented"
	return nil
}

type generateSchemaOverride struct {
	refEncoding map[jsonschema.Ref]ir.Encoding
	nameRef     func(ref jsonschema.Ref, def refNamer) (string, error)
	fieldMut    func(*ir.Field) error
	// request indicates this schema is for a request body (not response).
	// Used to decide how to handle empty schemas.
	request bool
}

func (g *Generator) generateSchema(
	ctx *genctx,
	name string,
	schema *jsonschema.Schema,
	optional bool,
	override *generateSchemaOverride,
) (_ *ir.Type, rerr error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GenerateSchemaOptions is options structure for GenerateSchema.
type GenerateSchemaOptions struct {
	// TypeName is root schema type name. Defaults to "Type".
	TypeName string
	// FileName is output filename. Defaults to "output.gen.go".
	FileName string
	// PkgName is the package name. Defaults to GOPACKAGE environment variable, if any. Otherwise, to "output".
	PkgName string
	// TrimPrefix is a ref name prefixes to trim. Defaults to []string{"#/definitions/", "#/$defs/"}.
	TrimPrefix []string
	// Logger to use.
	Logger *zap.Logger
}

func (o *GenerateSchemaOptions) setDefaults() { _ = "STUB: not implemented"; return }

// GenerateSchema generates type, validation and JSON encoding for given schema.
func GenerateSchema(schema *jsonschema.Schema, fs FileSystem, opts GenerateSchemaOptions) (rerr error) {
	_ = "STUB: not implemented"
	return nil
}

// TODO(tdakkota): pass input filename
