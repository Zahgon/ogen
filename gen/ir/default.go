package ir

// Default represents default value.
type Default struct {
	Value any
	Set   bool
}

// IsNil whether value is set, but null.
func (d Default) IsNil() bool { _ = "STUB: not implemented"; return false }
