package ir

// ExternalType represents an external type.
type ExternalType struct {
	PackagePath string
	PackageName string
	ImportAlias string
	TypeName    string
	Encode      ExternalEncoding
	Decode      ExternalEncoding
	IsPointer   bool
}

// String returns the string representation of the ExternalType.
func (e ExternalType) Primitive() PrimitiveType {
	_ = "STUB: not implemented"
	return *new(PrimitiveType)
}

// ExternalEncoding is a kind of external type.
type ExternalEncoding int

// String returns the string representation of the ExternalEncoding.
func (e ExternalEncoding) String() string { _ = "STUB: not implemented"; return "" }

// Has checks if the encoding kind is present in the ExternalEncoding.
func (e ExternalEncoding) Has(kind ExternalEncoding) bool { _ = "STUB: not implemented"; return false }

const (
	// ExternalNative indicates that the type implements ogen's json.Marshaler or json.Unmarshaler.
	ExternalNative ExternalEncoding = 1 << iota
	// ExternalJSON indicates that the type implements stdlib json.Marshaler or json.Unmarshaler.
	ExternalJSON
	// ExternalText indicates that the type implements stdlib encoding.TextMarshaler or encoding.TextUnmarshaler.
	ExternalText
	// ExternalBinary indicates that the type implements stdlib encoding.BinaryMarshaler or encoding.BinaryUnmarshaler.
	ExternalBinary
)

func getExternalType(input string) (ExternalType, error) {
	_ = "STUB: not implemented"
	return *new(ExternalType), nil
}

func parseTypePath(input string) (pkgPath, typeName string, isPointer bool, _ error) {
	_ = "STUB: not implemented"
	return "", "", false, nil
}

// Check for pointer prefix

// Parse package path

// Look for matching ')'

// skip ')'

// skip '.'

// No parens, assume last '.' separates package and type

// adjust for initial offset

var encoders = map[[2]string]ExternalEncoding{
	{"github.com/ogen-go/ogen/json", "Marshaler"}: ExternalNative,
	{"encoding/json", "Marshaler"}:                ExternalJSON,
	{"encoding", "TextMarshaler"}:                 ExternalText,
	{"encoding", "BinaryMarshaler"}:               ExternalBinary,
}

var decoders = map[[2]string]ExternalEncoding{
	{"github.com/ogen-go/ogen/json", "Unmarshaler"}: ExternalNative,
	{"encoding/json", "Unmarshaler"}:                ExternalJSON,
	{"encoding", "TextUnmarshaler"}:                 ExternalText,
	{"encoding", "BinaryUnmarshaler"}:               ExternalBinary,
}

func loadExternal(pkgPath, typeName string) (pkgName string, encode, decode ExternalEncoding, _ error) {
	_ = "STUB: not implemented"
	return "", *new(ExternalEncoding), *new(ExternalEncoding), nil
}
