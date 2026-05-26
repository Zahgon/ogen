package ir

// JSON returns json encoding/decoding rules for t.
func (t *Type) JSON() JSON {
	_ = "STUB: not implemented"
	return *

	// JSON specifies json encoding and decoding for Type.
	new(JSON)
}

type JSON struct {
	t      *Type
	except []string
}

// AnyFields whether if type has any fields to encode.
func (j JSON) AnyFields() bool { _ = "STUB: not implemented"; return false }

// NeedsReceiver reports whether encoding the fields would reference the
// receiver variable "s".  This is false when every non-excluded, non-inline
// field carries a const value (encoded as a literal) and there are no
// additional-properties, pattern-properties or inline-sum fields.
func (j JSON) NeedsReceiver() bool { _ = "STUB: not implemented"; return false }

// Inline fields (additional / pattern / sum props) always reference s.

// A non-const regular field will be encoded via field_elem → s.Name.

// Except return JSON with filter by given properties.
func (j JSON) Except(set ...string) JSON { _ = "STUB: not implemented"; return *new(JSON) }

type JSONFields []*Field

// FirstRequiredIndex returns first required field index.
//
// Or -1 if there is no required fields.
func (j JSONFields) FirstRequiredIndex() int { _ = "STUB: not implemented"; return 0 }

// HasRequired whether object has required fields
func (j JSONFields) HasRequired() bool { _ = "STUB: not implemented"; return false }

// RequiredMask returns array of 64-bit bitmasks for required fields.
func (j JSONFields) RequiredMask() []uint8 { _ = "STUB: not implemented"; return nil }

// Fields return all fields of Type that should be encoded via json.
func (j JSON) Fields() (fields JSONFields) { _ = "STUB: not implemented"; return *new(JSONFields) }

// AdditionalProps return field of Type that should be encoded as inlined map.
func (j JSON) AdditionalProps() *Field { _ = "STUB: not implemented"; return nil }

// PatternProps return field of Type that should be encoded as inlined map with pattern.
func (j JSON) PatternProps() (fields []*Field) { _ = "STUB: not implemented"; return nil }

// SumProps return field of Type that should be encoded as inlined sum.
func (j JSON) SumProps() (fields []*Field) { _ = "STUB: not implemented"; return nil }

// Format returns format name for handling json encoding or decoding.
//
// Mostly used for encoding or decoding of string formats, like `json.EncodeUUID`,
// where UUID is Format.
func (j JSON) Format() string { _ = "STUB: not implemented"; return "" }

// Type returns json value type that can represent Type.
//
// E.g. string primitive can be represented by StringValue which is commonly
// returned from `i.WhatIsNext()` method.
// Blank string is returned if there is no appropriate json type.
func (j JSON) Type() string { _ = "STUB: not implemented"; return "" }

func collectTypes(t *Type, types map[string]struct{}) { _ = "STUB: not implemented"; return }

// SumTypes returns jx.Type list for this sum type.
func (j JSON) SumTypes() string { _ = "STUB: not implemented"; return "" }

const arraySuffix = "Array"

func jsonType(t *Type) string { _ = "STUB: not implemented"; return "" }

// raw denotes whether Type can be encoded or decoded using simple
// json method, e.g. j.WriteString.
//
// Mostly true for primitives or enums.
func (j JSON) raw() bool { _ = "STUB: not implemented"; return false }

func (j JSON) Decode() string {
	_ = "STUB: not implemented"

	// Copy to prevent referencing internal buffer.
	return ""
}

// No arguments.

// Fn returns jx.Encoder or jx.Decoder method name.
//
// If blank, value cannot be encoded with single method call.
func (j JSON) Fn() string { _ = "STUB: not implemented"; return "" }

// IsBase64 whether field has base64 encoding.
func (j JSON) IsBase64() bool { _ = "STUB: not implemented"; return false }

// TimeFormat returns time format for json encoding and decoding.
func (j JSON) TimeFormat() string { _ = "STUB: not implemented"; return "" }

// Encoder returns format name for handling json encoding.
//
// Mostly used for encoding of string formats, like `json.EncodeUUID`, where
// UUID is Encoder.
func (j JSON) Encoder() string { _ = "STUB: not implemented"; return "" }

// Decoder returns format name for handling json decoding.
//
// Mostly used for decoding of string formats, like `json.DecodeUUID`, where
// UUID is Decoder.
func (j JSON) Decoder() string { _ = "STUB: not implemented"; return "" }

// Sum returns specification for parsing value as sum type.
func (j JSON) Sum() SumJSON { _ = "STUB: not implemented"; return *new(SumJSON) }

// Check for field-based discrimination (UniqueFields or ValueDiscriminators on sum type)

// Check for unique fields on variants (legacy approach)

type SumJSONType byte

const (
	SumJSONPrimitive SumJSONType = iota
	SumJSONFields
	SumJSONDiscriminator
	SumJSONTypeDiscriminator
)

// SumJSON specifies rules for parsing sum types in json.
type SumJSON struct {
	Type SumJSONType
}

func (s SumJSON) String() string { _ = "STUB: not implemented"; return "" }

func (s SumJSON) Primitive() bool         { _ = "STUB: not implemented"; return false }
func (s SumJSON) Discriminator() bool     { _ = "STUB: not implemented"; return false }
func (s SumJSON) TypeDiscriminator() bool { _ = "STUB: not implemented"; return false }
func (s SumJSON) Fields() bool            { _ = "STUB: not implemented"; return false }
