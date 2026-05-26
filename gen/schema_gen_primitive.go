package gen

import (
	"github.com/ogen-go/ogen/gen/ir"
	"github.com/ogen-go/ogen/jsonschema"
)

func (g *schemaGen) primitive(name string, schema *jsonschema.Schema) (*ir.Type, error) {
	_ = "STUB: not implemented"
	return nil,

		// If const is set, treat it as a const value (not enum)
		// Const takes precedence over enum
		nil
}

func (g *schemaGen) enum(name string, t *ir.Type, schema *jsonschema.Schema) (*ir.Type, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// We accept 2 types of enums: ints and strings. However, for formatted
// string enums, we don't want to allow time/date/date-time formats as they
// require special handling

// Reject time-related formats for string enums until we properly handle them

func (g *schemaGen) validateEnumValues(s *jsonschema.Schema) error {
	_ = "STUB: not implemented"
	return nil
}

func (g *schemaGen) parseSimple(schema *jsonschema.Schema) *ir.Type {
	_ = "STUB: not implemented"
	return nil
}

// TODO(tdakkota): check ContentEncoding field

// Fallback to default.

func TypeFormatMapping() map[jsonschema.SchemaType]map[string]ir.PrimitiveType {
	_ = "STUB: not implemented"
	return nil
}

// FIXME(tdakkota): add decoder for int8, int16, uint8, uint16 to jx.

// See https://github.com/ogen-go/ogen/issues/307.

// Custom format, see https://github.com/ogen-go/ogen/issues/309.

// See https://github.com/ogen-go/ogen/issues/307.

// See https://github.com/ogen-go/ogen/issues/957.
