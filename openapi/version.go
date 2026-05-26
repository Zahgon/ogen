package openapi

// Version represents OpenAPI version.
type Version struct {
	// Major is the major version number.
	Major int
	// Minor is the minor version number.
	Minor int
	// Patch is the patch version number.
	Patch int
}

// MarshalText implements encoding.TextMarshaler.
func (v *Version) MarshalText() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// UnmarshalText implements encoding.TextUnmarshaler.
func (v *Version) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

// String returns the string representation of the version.
func (v Version) String() string { _ = "STUB: not implemented"; return "" }
