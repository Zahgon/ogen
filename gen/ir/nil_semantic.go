package ir

// NilSemantic specifies nil value semantics.
type NilSemantic string

// Possible nil value semantics.
const (
	NilInvalid  NilSemantic = "invalid"  // nil is invalid
	NilOptional NilSemantic = "optional" // nil is "no value"
	NilNull     NilSemantic = "null"     // nil is null
)

func (n NilSemantic) Invalid() bool  { _ = "STUB: not implemented"; return false }
func (n NilSemantic) Optional() bool { _ = "STUB: not implemented"; return false }
func (n NilSemantic) Null() bool     { _ = "STUB: not implemented"; return false }
