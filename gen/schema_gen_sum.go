package gen

import (
	"github.com/ogen-go/ogen/gen/ir"
	"github.com/ogen-go/ogen/jsonschema"
)

// fieldSignature represents a field's discriminating characteristics (name + type).
//
// This enables type-based field discrimination: fields with the same name but different
// types are not considered "common" and can be used for discrimination.
//
// For example, if VariantA has {id: string} and VariantB has {id: integer}, the "id"
// field can help discriminate between them even though the field names are identical.
type fieldSignature struct {
	name   string
	typeID string
}

const (
	typeIDAny     = "any"
	typeIDBoolean = "boolean"
	typeIDInteger = "integer"
	typeIDNumber  = "number"
	typeIDString  = "string"
	typeIDNull    = "null"
	typeIDObject  = "object"
	typeIDSum     = "sum"
	typeIDAlias   = "alias"
	typeIDPointer = "pointer"

	// jxTypeArray is the string representation of jx.Array for template generation.
	jxTypeArray = "jx.Array"
)

// jxTypeForFieldType returns the jx.Type constant name for runtime type checking.
// Returns empty string if the type is not distinguishable at JSON level.
func jxTypeForFieldType(ft *ir.Type) string { _ = "STUB: not implemented"; return "" }

// TODO(tdakkota): properly figure out JSON type.

// Enums serialize as strings in JSON

// getArrayElementTypeInfo extracts element type information from an array type ID.
// Returns the element type ID and its corresponding jx.Type.
// For non-array types, returns empty strings.
func getArrayElementTypeInfo(t *ir.Type) (elementTypeID, elementJxType string) {
	_ = "STUB: not implemented"
	return "", ""
}

// Extract element type: "array[string]" -> "string"

// Get the jx.Type for the element

// getFieldTypeID returns a type identifier for discrimination purposes.
// Fields with the same name but different typeIDs can discriminate variants.
func getFieldTypeID(t *ir.Type) string { _ = "STUB: not implemented"; return "" }

// Unwrap optionals and nullables to get the base type

// Enums are distinct from their underlying types

func canUseTypeDiscriminator(sum []*ir.Type, isOneOf bool) bool {
	_ = "STUB: not implemented"

	// Collect map of variant kinds.
	return false
}

// Special case for anyOf with integer and number.

// Cannot make type discriminator with Any.

// Type kind is not unique, so we cannot distinguish variants by type.

// TODO(tdakkota): Do not allow type discriminator for nested sum types with integer and
// 	number variants at the same time. We can add support for this later, but it's not trivial.

func ensureNoInfiniteRecursion(parent *jsonschema.Schema) error {
	_ = "STUB: not implemented"
	return nil
}

// Just skip nil schemas. We handle them later.

func (g *schemaGen) collectSumVariants(
	name string,
	schemas []*jsonschema.Schema,
) (sum []*ir.Type, _ error) {
	_ = "STUB: not implemented"
	// TODO(tdakkota): convert oneOf+null into generic
	return nil, nil
}

// generate without boxing because:
// 1) sum variant cannot be optional
// 2) if sum variant is nullable - null type already added into sum

func schemaName(k jsonschema.Ref) (string, bool) { _ = "STUB: not implemented"; return "", false }

// handleExplicitDiscriminator processes explicit discriminator mappings for both oneOf and anyOf.
// Returns true if discriminator was handled, false if no discriminator present.
func (g *schemaGen) handleExplicitDiscriminator(sum *ir.Type, schema *jsonschema.Schema, variants []*jsonschema.Schema) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// Build mappings and collect keys

// Only proceed if we have explicit mappings

// Validate: Check if discriminator field uses value-based discrimination
// Collect the discriminator field type from each variant
// variant name -> jxType

// Find the discriminator field in this variant

// Check if all discriminator fields have the same empty jxType (value-based discrimination)

// Set discriminator only if we have mappings

// Generate names using the helper

// Generate Go variable name for each mapping Key

func (g *schemaGen) anyOf(name string, schema *jsonschema.Schema, side bool) (*ir.Type, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Here we try to create sum type from anyOf for variants with JSON type-based discriminator.

// Check for explicit discriminator

func (g *schemaGen) oneOf(name string, schema *jsonschema.Schema, side bool) (*ir.Type, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Collect variants first to enable early validation

// Early validation for explicit discriminator (before regtype)
// This prevents broken types from being registered when discriminator is invalid

// Quick validation: check if this would be value-based discrimination

// Only validate if there are explicit mappings

// Check if all discriminator fields have same empty jxType (value-based)

// 1st case: explicit discriminator.

// 2nd case: implicit mapping based on schema references (only if discriminator present but no explicit mappings).

// Implicit mapping, defaults to type name.

// Spec says (https://spec.openapis.org/oas/v3.1.0#discriminator-object):
//
// 	The expectation now is that a property with name petType MUST be present in the response payload,
// 	and the value will correspond to the name of a schema defined in the OAS document
//
// What is name of a schema? Is it the last part of the pointer?
// What if pointer part of reference is empty, like `User.json#`?
//
// As always, OpenAPI is not clear enough.

// Name will be set below

// Generate names using the helper (only if we have mappings)

// Set discriminator only if we have mappings

// Apply generated names to mappings

// 3rd case: distinguish by serialization type.

// 4th case: distinguish by unique fields (considering field types).

// Determine unique fields for each SumOf variant.
// We now track field signatures (name + type) instead of just names.

// variant -> fieldName -> signatures

// Collect field signatures that are common to at least 2 variants.
// A field signature is common if it appears in multiple variants with the exact same name AND type.

// variant and otherVariant have common field signature.
// This means same field name with same type.

// Delete common field signatures from unique sets.

// Check that at most one type has no unique fields.

// Set mapping without unique fields as default

// Unable to deterministically select sub-schema only on fields.

// Collect field -> variant mapping to compute fields used by multiple variants.

// Collect the problematic variants and fields.

// Build sorted list of variants with their unique field names.
// Use alphabetical sorting to match upstream behavior and minimize diffs.

// Extract unique field names from signatures

// Sort field names alphabetically

// Sort variants alphabetically by name

// Populate SumSpec.Unique with fields that are discriminating.
// Also build UniqueFields map for template iteration.

// Initialize UniqueFieldTypes map for runtime type checking

// Iterate through fields in schema order

// Check if this field name is in the unique list

// Verify the field signature is actually unique

// Check if this field was already added to Unique

// Store expected jx.Type for runtime type checking

// Check if field is nullable (can be null in JSON)
// A field is nullable if:
// 1. It's a generic type (KindGeneric) with Nullable=true (e.g., NilString, OptNilInt)
// 2. It's a pointer with Null semantic

// Get array element type info for array element discrimination

// Add to UniqueFields map for template iteration
// Include entries even when jxType is empty (simple field-name discrimination)

// Empty string means no runtime type check needed
// true if field accepts null values

// canUseValueDiscrimination checks if a field can discriminate variants by enum values.
// Returns true if all variants have non-overlapping enum values for this field.

// Collect enum values for each variant

// set of enum values

// Find the IR type for this variant

// Find the field in this variant

// Unwrap optionals/pointers to get base type

// Check if it's an enum

// Not an enum, can't use value discrimination

// Collect enum values

// EnumVariant.Value is the actual value (string, int, etc.)
// For string enums, it's already a string

// Non-string enums (shouldn't happen for JSON strings, but be safe)

// Empty enum, can't discriminate

// Check that we found enum info for all variants

// Some variants don't have enum types

// Check for overlapping enum values

// Find overlaps

// Build error message about overlapping values

// Build the value -> variant mapping

// Validate that we can actually discriminate variants after jxType deduplication
// This catches cases like arrays with different element types that both map to jx.Array
//
// We need to find at least one field that can discriminate all variants.
// If a field has overlapping enum values, we skip it and try other fields.
// Only fail if NO field can discriminate.

// Track fields that need discrimination but couldn't be discriminated

// Single variant, no need to discriminate

// Count unique jxTypes for this field

// If all variants have the same jxType (or empty), try value-based or array element discrimination

// Try value-based discrimination (enum values)

// Overlapping enum values - record this but continue checking other fields

// Try next field

// Initialize map if needed

// Store the value discriminator

// Remove from UniqueFields to avoid duplicate case statements in template

// This field can discriminate, move to next field

// Value discrimination didn't work, check if array element discrimination is possible

// If all variants are arrays with different element types, check if we have other discriminating fields

// Array element discrimination works, but we need to check if there are other
// unique fields to discriminate in case this array field is missing or empty.
// If this array field is the ONLY way to discriminate, we should reject it
// because the field might be optional and missing from the JSON.

// Check if any variant has a unique field that exists ONLY in that variant (by name)

// variant -> set of field names

// For each variant, check if it has any field name unique to it

// Can discriminate by array element type (with fallback to other fields)

// Fall through to the error below - array element discrimination alone is not sufficient

// Can't use value discrimination or array element discrimination, record for potential error

// If we have undiscriminable fields and no successful discriminator was found,
// we need to fail. But if at least one field can discriminate, we're okay.

// Use the first undiscriminable field for the error message

func (g *schemaGen) allOf(name string, schema *jsonschema.Schema) (*ir.Type, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// shallowSchemaCopy returns a shallow copy of the given schema.
//
// If given schema is nil, nil is returned.
//
// All references in Schema are shallow copied.
func shallowSchemaCopy(s *jsonschema.Schema) *jsonschema.Schema {
	_ = "STUB: not implemented"
	return nil
}

func flattenAllOfSchema(schema *jsonschema.Schema) (*jsonschema.Schema, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// If there is only one schema in allOf, avoid merging and keep the inner schema
// while still applying wrapper-level metadata from the parent.

// The reference field must not change.

func mergeNSchemes(ss []*jsonschema.Schema) (_ *jsonschema.Schema, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func mergeSchemes(s1, s2 *jsonschema.Schema) (_ *jsonschema.Schema, err error) {
	_ = "STUB: not implemented"
	// Helper functions for comparing validation fields.
	return nil, nil
}

// TODO(tdakkota): handle in a better way.

// Type

// Format

// Enum

// Default

// Nothing to do.

// Const

// Discriminator

// TODO(tdakkota): implement

// String validation

// Integer, Number validation

// NOTE: We need to refactor ir.Validators to support multiple 'multipleOf's.
//
// Most likely it will require rewriting this schema merging code, because
// we cannot set multiple 'multipleOf's into single jsonschema.Schema.
// We need to generate ir.Type for each schema in 'allOf' and then merge
// them into single *ir.Type with all the validation.

// Array validation

// Object validation

// Nothing to do.

// oneOf, anyOf

// mergeProperties finds properties with identical names
// and tries to merge them into one, avoiding duplicates.
func mergeProperties(s1, s2 *jsonschema.Schema) ([]jsonschema.Property, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Fill the map with p1 props.

// Try to merge p2 props.

// Property name conflict.

// TODO(tdakkota): handle in a better way.

func mergeEnums(s1, s2 *jsonschema.Schema) ([]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Keep values that are present in both enums.

// FIXME(tdakkota): quadratic complexity.
