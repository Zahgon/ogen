package ir

func (t *Type) EncodeFn() string { _ = "STUB: not implemented"; return "" }

func (t *Type) IsBase64Stream() bool { _ = "STUB: not implemented"; return false }

func (t Type) uriFormat() string { _ = "STUB: not implemented"; return "" }

// externalType determines the most appropriate external encoding interface to use.
func (t Type) externalType(e, prefer ExternalEncoding) ExternalEncoding {
	_ = "STUB: not implemented"
	return *new(ExternalEncoding)
}

func (t Type) ToString() string { _ = "STUB: not implemented"; return "" }

func (t Type) FromString() string { _ = "STUB: not implemented"; return "" }

func (t *Type) IsInteger() bool { _ = "STUB: not implemented"; return false }

func (t *Type) IsFloat() bool { _ = "STUB: not implemented"; return false }

func (t *Type) IsDecimal() bool { _ = "STUB: not implemented"; return false }

func (t *Type) IsStringifiedFloat() bool { _ = "STUB: not implemented"; return false }

func (t *Type) IsNull() bool { _ = "STUB: not implemented"; return false }

func (t *Type) IsArray() bool     { _ = "STUB: not implemented"; return false }
func (t *Type) IsMap() bool       { _ = "STUB: not implemented"; return false }
func (t *Type) IsPrimitive() bool { _ = "STUB: not implemented"; return false }
func (t *Type) IsStruct() bool    { _ = "STUB: not implemented"; return false }
func (t *Type) IsPointer() bool   { _ = "STUB: not implemented"; return false }
func (t *Type) IsEnum() bool      { _ = "STUB: not implemented"; return false }
func (t *Type) IsGeneric() bool   { _ = "STUB: not implemented"; return false }
func (t *Type) IsAlias() bool     { _ = "STUB: not implemented"; return false }
func (t *Type) IsInterface() bool { _ = "STUB: not implemented"; return false }
func (t *Type) IsSum() bool       { _ = "STUB: not implemented"; return false }
func (t *Type) IsAny() bool       { _ = "STUB: not implemented"; return false }
func (t *Type) IsStream() bool    { _ = "STUB: not implemented"; return false }
func (t *Type) IsNumeric() bool   { _ = "STUB: not implemented"; return false }
func (t *Type) IsExternal() bool  { _ = "STUB: not implemented"; return false }

func (t *Type) MustField(name string) *Field { _ = "STUB: not implemented"; return nil }

func (t *Type) SetFieldType(name string, newT *Type) { _ = "STUB: not implemented"; return }

// TypeDiscriminatorCase is a helper struct for describing type discriminator case.
type TypeDiscriminatorCase struct {
	// JXTypes is jx.Type values list.
	JXTypes string
	// Type is the type to be used for this case.
	Type *Type
	// IntType is the type to be used for this case when the type discriminator should distinguish
	// between integer and float types.
	IntType *Type
}

func (t *Type) TypeDiscriminator() (r []TypeDiscriminatorCase) {
	_ = "STUB: not implemented"
	return nil
}

// DoPassByPointer returns true if type should be passed by pointer.
func (t *Type) DoPassByPointer() bool { _ = "STUB: not implemented"; return false }

// ReadOnlyReceiver returns the receiver type for read-only methods.
func (t *Type) ReadOnlyReceiver() string { _ = "STUB: not implemented"; return "" }
