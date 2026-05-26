// Package location provides utilities to track values over the spec.
package location

// Pointer is a location of a value.
type Pointer struct {
	// Source is the File where the value is located.
	Source File
	// Locator stores the Position of a value.
	Locator Locator
}

// File returns the File where the value is located.
func (p Pointer) File() File {
	_ = "STUB: not implemented"

	// Position returns the position of the value if it is set.
	return *new(File)
}

func (p Pointer) Position() (Position, bool) {
	_ = "STUB: not implemented"
	return *new(Position), false
}

// Key tries to find the child node using given key and returns its pointer.
//
// See Key method of Locator.
func (p Pointer) Key(key string) (ptr Pointer) { _ = "STUB: not implemented"; return *new(Pointer) }

// Field tries to find the child node using given key and returns its pointer.
//
// See Field method of Locator.
func (p Pointer) Field(key string) (ptr Pointer) { _ = "STUB: not implemented"; return *new(Pointer) }

// Index tries to find the child node using given index and returns its pointer.
//
// See Index method of Locator.
func (p Pointer) Index(idx int) (ptr Pointer) { _ = "STUB: not implemented"; return *new(Pointer) }
