package ir

// PrimitiveType represents a primitive type in Go.
type PrimitiveType string

func (p PrimitiveType) String() string {
	_ = "STUB: not implemented"

	// IsString whether this type is string.
	return ""
}

func (p PrimitiveType) IsString() bool {
	_ = "STUB: not implemented"

	// Primitive types.
	return false
}

const (
	None      PrimitiveType = ""
	String    PrimitiveType = "string"
	ByteSlice PrimitiveType = "[]byte"
	Int       PrimitiveType = "int"
	Int8      PrimitiveType = "int8"
	Int16     PrimitiveType = "int16"
	Int32     PrimitiveType = "int32"
	Int64     PrimitiveType = "int64"
	Uint      PrimitiveType = "uint"
	Uint8     PrimitiveType = "uint8"
	Uint16    PrimitiveType = "uint16"
	Uint32    PrimitiveType = "uint32"
	Uint64    PrimitiveType = "uint64"
	Float32   PrimitiveType = "float32"
	Float64   PrimitiveType = "float64"
	Bool      PrimitiveType = "bool"
	Null      PrimitiveType = "struct{}"
	Time      PrimitiveType = "time.Time"
	Duration  PrimitiveType = "time.Duration"
	UUID      PrimitiveType = "uuid.UUID"
	MAC       PrimitiveType = "net.HardwareAddr"
	IP        PrimitiveType = "netip.Addr"
	URL       PrimitiveType = "url.URL"
	File      PrimitiveType = "ht.MultipartFile"
	Decimal   PrimitiveType = "decimal.Decimal"
	Custom    PrimitiveType = "<custom>"
)
