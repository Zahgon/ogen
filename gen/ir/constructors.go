package ir

import (
	"github.com/ogen-go/ogen/jsonschema"
)

func Primitive(typ PrimitiveType, schema *jsonschema.Schema) *Type {
	_ = "STUB: not implemented"
	return nil
}

func Array(item *Type, sem NilSemantic, schema *jsonschema.Schema) *Type {
	_ = "STUB: not implemented"
	return nil
}

func Alias(name string, to *Type) *Type { _ = "STUB: not implemented"; return nil }

func Interface(name string) *Type { _ = "STUB: not implemented"; return nil }

func Pointer(to *Type, sem NilSemantic) *Type { _ = "STUB: not implemented"; return nil }

func Generic(name string, of *Type, v GenericVariant) *Type { _ = "STUB: not implemented"; return nil }

func Any(schema *jsonschema.Schema) *Type { _ = "STUB: not implemented"; return nil }

func Stream(name string, schema *jsonschema.Schema) *Type { _ = "STUB: not implemented"; return nil }

func External(schema *jsonschema.Schema) (*Type, error) {
	_ = "STUB: not implemented"
	// If schema.XOgenType has no slashes or dots, it is a builtin type.
	return nil, nil
}
