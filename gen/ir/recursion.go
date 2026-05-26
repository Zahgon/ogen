package ir

func (t *Type) RecursiveTo(target *Type) bool { _ = "STUB: not implemented"; return false }

func (t *Type) recursive(target *Type, path *walkpath) bool {
	_ = "STUB: not implemented"

	// This is a list of types that cannot cause recursion.
	//
	// Primitive - has no fields.
	// Enum      - has no fields.
	// Any       - has no fields.
	// Pointer   - prevents recursion.
	// Array     - prevents recursion.
	// Map       - prevents recursion.
	return false
}

type walkpath struct {
	nodes map[*Type]struct{}
}

func (wp *walkpath) has(t *Type) bool { _ = "STUB: not implemented"; return false }

func (wp *walkpath) add(t *Type) { _ = "STUB: not implemented"; return }

func (wp *walkpath) delete(t *Type) { _ = "STUB: not implemented"; return }
