package gen

import (
	"go.uber.org/zap"

	"github.com/ogen-go/ogen/gen/ir"
	"github.com/ogen-go/ogen/jsonschema"
)

type refNamer func(ref jsonschema.Ref) (string, error)

const defaultSchemaDepthLimit = 1000

type schemaGen struct {
	side      []*ir.Type
	localRefs map[jsonschema.Ref]*ir.Type
	lookupRef func(ref jsonschema.Ref) (*ir.Type, bool)
	nameRef   func(ref jsonschema.Ref) (string, error)
	fieldMut  func(*ir.Field) error
	fail      func(err error) error
	imports   map[string]string

	depthLimit int
	depthCount int

	request bool // true if generating for request body

	log *zap.Logger
}

func newSchemaGen(lookupRef func(ref jsonschema.Ref) (*ir.Type, bool)) *schemaGen {
	_ = "STUB: not implemented"
	return nil
}

func variantFieldName(t *ir.Type) string { _ = "STUB: not implemented"; return "" }

type schemaDepthError struct {
	limit int
}

func (e *schemaDepthError) Error() string { _ = "STUB: not implemented"; return "" }

func handleSchemaDepth(s *jsonschema.Schema, rerr *error) { _ = "STUB: not implemented"; return }

// Ensure that schema is not nil.

// Try to use location.Error.

func (g *schemaGen) generate(name string, schema *jsonschema.Schema, optional bool) (*ir.Type, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Panicing is not cool, but is better rather than wrap the error N = depthLimit times.

// TODO: update refcache to point to new boxed type?

func (g *schemaGen) generate2(name string, schema *jsonschema.Schema) (ret *ir.Type, err error) {
	_ = "STUB: not implemented"

	// Empty schema (no schema field in OpenAPI spec).
	// For responses: Allow jx.Raw since client must handle unknown JSON from server.
	// For requests: Reject to avoid clients sending arbitrary data without spec guidance.
	return nil, nil
}

// For responses, treat as "any valid JSON value" (jx.Raw).
// Consistent with array item handling (line 437).

// Do not fail schema generation if we cannot handle defaults.

// Only error on truly invalid cases (nil or empty item type).
// Complex types (Array, Object) are now supported via Equal/Hash generation.

// Primitive enums are handled below

// Non-primitive object enums generate sum types with struct variants.
// Each enum value becomes a concrete struct type.

// Array enums and empty type enums are treated as "any" type.
// The enum constraint is documented in OpenAPI but not enforced at runtime.

// Stores spec name of the field.

// Use existing as a fallback.

// Create special field for additionalProperties.

// Using the name from the ogen schema extension.
// Avoiding name conflicts is up to user.

func (g *schemaGen) regtype(name string, t *ir.Type) *ir.Type {
	_ = "STUB: not implemented"
	return nil
}

func (g *schemaGen) checkDefaultType(s *jsonschema.Schema, val any) error {
	_ = "STUB: not implemented"

	// Schema has no validators.
	return nil
}

// nonPrimitiveObjectEnum generates a sum type for object enums.
// Each enum value becomes a concrete struct variant.
func (g *schemaGen) nonPrimitiveObjectEnum(name string, schema *jsonschema.Schema) (*ir.Type, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Convert enum values to map[string]any

// Find a discriminating field - a string field with unique values across all variants

// No discriminator found, fall back to index-based naming

// Create the sum type

// Generate struct types for each enum value

// Generate the variant struct type

// Add fields from the object

// Register and add to variants

// Set up discrimination

// Value-based discrimination on the discriminator field

// No discriminator field found, use type-based discrimination as fallback

// findEnumDiscriminator finds a string field that has unique values across all enum objects.
func findEnumDiscriminator(objects []map[string]any) (fieldName string, values []string) {
	_ = "STUB: not implemented"
	return "", nil
}

// Find all string fields present in all objects

// Find a field with unique values across all objects

// Field not present in all objects

// Check if all values are unique

// Use these values as variant names (capitalized)

// inferSchemaFromObject creates a jsonschema.Schema from an object literal.
func inferSchemaFromObject(obj map[string]any) *jsonschema.Schema {
	_ = "STUB: not implemented"
	return nil
}

// Sort keys for deterministic output

// inferSchemaFromValue creates a jsonschema.Schema from a JSON value.
func inferSchemaFromValue(v any) *jsonschema.Schema { _ = "STUB: not implemented"; return nil }

// inferTypeFromValue creates an ir.Type from a JSON value.
func (g *schemaGen) inferTypeFromValue(v any, schema *jsonschema.Schema) *ir.Type {
	_ = "STUB: not implemented"
	return nil
}
