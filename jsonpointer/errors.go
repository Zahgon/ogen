package jsonpointer

// NotFoundError reports that requested value is not found.
type NotFoundError struct {
	Pointer string
}

// Error implements error.
func (n *NotFoundError) Error() string { _ = "STUB: not implemented"; return "" }
