package ir

import (
	"github.com/ogen-go/ogen/jsonschema"
	"github.com/ogen-go/ogen/ogenregex"
)

type Kind string

const (
	KindPrimitive Kind = "primitive"
	KindArray     Kind = "array"
	KindMap       Kind = "map"
	KindAlias     Kind = "alias"
	KindConst     Kind = "const"
	KindEnum      Kind = "enum"
	KindStruct    Kind = "struct"
	KindPointer   Kind = "pointer"
	KindInterface Kind = "interface"
	KindGeneric   Kind = "generic"
	KindSum       Kind = "sum"
	KindAny       Kind = "any"
	KindStream    Kind = "stream"
)

type SumSpecMap struct {
	Key               string
	Type              *Type
	DiscriminatorType *Type
	Name              string
}

// UniqueFieldVariant represents a variant that has a specific unique field.
type UniqueFieldVariant struct {
	VariantName string // e.g., "SystemEvent"
	VariantType string // e.g., "SystemEventEvent"
	FieldType   string // jx.Type constant, e.g., "jx.String"
	Nullable    bool   // true if field is nullable (accepts both base type and jx.Null)

	// ArrayElementType is the jx.Type of array elements for array element discrimination.
	// Only set when FieldType is "jx.Array" and element type can distinguish variants.
	// e.g., "jx.String" for array[string], "jx.Number" for array[integer], "jx.Object" for array[object]
	ArrayElementType string

	// ArrayElementTypeID is the full type ID for array elements (e.g., "string", "integer", "object").
	// Used for more detailed discrimination like distinguishing integer vs number.
	ArrayElementTypeID string
}

// SumSpec for KindSum.
type SumSpec struct {
	Unique []*Field
	// DefaultMapping is name of default mapping.
	//
	// Used for variant which has no unique fields.
	DefaultMapping string

	// Discriminator is field name of sum type discriminator.
	Discriminator string
	// Mapping is discriminator value -> variant mapping.
	Mapping []SumSpecMap

	// TypeDiscriminator denotes to distinguish variants by type.
	TypeDiscriminator bool

	// UniqueFieldTypes maps field JSON names to their expected jx.Type for type-based discrimination.
	// Key: field JSON name, Value: jx.Type constant name (e.g., "jx.String", "jx.Number")
	// Only populated for fields that require runtime type checking.
	UniqueFieldTypes map[string]string

	// UniqueFields maps field names to variants that have that field as unique.
	// Used for generating field-based discrimination in oneOf/anyOf.
	// Key: field JSON name, Value: list of variants with that unique field
	UniqueFields map[string][]UniqueFieldVariant

	// ValueDiscriminators maps field names to value-based discriminators.
	// Used when variants have the same field name and JSON type but different enum values.
	// Key: field JSON name, Value: ValueDiscriminator with enum value mappings
	ValueDiscriminators map[string]ValueDiscriminator
}

// ValueDiscriminator represents a field that discriminates variants by enum value.
type ValueDiscriminator struct {
	// FieldName is the JSON field name used for discrimination
	FieldName string
	// ValueToVariant maps enum values to variant type constants
	// Key: enum value (e.g., "active"), Value: variant type constant (e.g., "ActiveStatusResponse")
	ValueToVariant map[string]string
}

type ResolvedSumSpecMap struct {
	Name              string
	Key               string
	DiscriminatorType *Type
}

type PickedMappingEntries []*ResolvedSumSpecMap

func (e PickedMappingEntries) JoinConstNames() string { _ = "STUB: not implemented"; return "" }

// PickMappingEntriesFor returns all mapping entries for given type they exists.
func (s SumSpec) PickMappingEntriesFor(t, sumOf *Type) PickedMappingEntries {
	_ = "STUB: not implemented"
	return *new(PickedMappingEntries)
}

// Deprecated: use PickMappingEntriesFor instead.
//
// PickMappingEntryFor returns the first mapping entry for given type if exists.
func (s SumSpec) PickMappingEntryFor(t *Type) *SumSpecMap { _ = "STUB: not implemented"; return nil }

type Type struct {
	Doc                 string              // ogen documentation
	Kind                Kind                // kind
	Name                string              // only for struct, alias, interface, enum, stream, generic, map, sum
	Primitive           PrimitiveType       // only for primitive, enum
	AliasTo             *Type               // only for alias
	PointerTo           *Type               // only for pointer
	SumOf               []*Type             // only for sum
	SumSpec             SumSpec             // only for sum
	Item                *Type               // only for array, map
	EnumVariants        []*EnumVariant      // only for enum
	Fields              []*Field            // only for struct
	Implements          map[*Type]struct{}  // only for struct, alias, enum
	Implementations     map[*Type]struct{}  // only for interface
	InterfaceMethods    map[string]struct{} // only for interface
	Schema              *jsonschema.Schema  // for all kinds except pointer, interface. Can be nil.
	NilSemantic         NilSemantic         // only for pointer
	GenericOf           *Type               // only for generic
	GenericVariant      GenericVariant      // only for generic
	MapPattern          ogenregex.Regexp    // only for map
	DenyAdditionalProps bool                // only for map and struct
	AllowedProps        map[string]struct{} // only for map and struct
	External            ExternalType        // only for custom type
	Validators          Validators
	Tuple               bool // only for struct
	// Features contains a set of features the type must implement.
	// Available features: 'json', 'uri'.
	//
	// If some of these features are set, generator
	// generates additional encoding methods if needed.
	Features []string
}

// GoDoc returns type godoc.
func (t Type) GoDoc() []string { _ = "STUB: not implemented"; return nil }

// Default returns default value of this type, if it is set.
func (t Type) Default() Default { _ = "STUB: not implemented"; return *new(Default) }

func (t Type) String() string { _ = "STUB: not implemented"; return "" }

func (t *Type) Pointer(sem NilSemantic) *Type { _ = "STUB: not implemented"; return nil }

// Format denotes whether custom formatting for Type is required while encoding
// or decoding.
//
// TODO(ernado): can we use t.JSON here?
func (t *Type) Format() bool { _ = "STUB: not implemented"; return false }

func (t *Type) Is(vs ...Kind) bool { _ = "STUB: not implemented"; return false }

// Go returns valid Go type for this Type.
func (t *Type) Go() string { _ = "STUB: not implemented"; return "" }

// NamePostfix returns name postfix for optional wrapper.
func (t *Type) NamePostfix() string { _ = "STUB: not implemented"; return "" }

// If type is external and has XOgenName, use it as name postfix.
// This is to be able to work around name conflicts where multiple
// packages have a type with the same name.
