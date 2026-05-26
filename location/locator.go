package location

import "github.com/go-faster/yaml"

// Locatable is an interface for JSON value position store.
type Locatable interface {
	// SetPosition sets the position of the value.
	SetPosition(Position)

	// Position returns the position of the value if it is set.
	Position() (Position, bool)
}

// Locator is a Position holder.
//
// Basically, it is a simple wrapper around Position to
// embed it to spec types.
type Locator struct {
	position Position
	set      bool
}

// Pointer makes a Pointer from the Locator and given File.
func (l Locator) Pointer(file File) Pointer { _ = "STUB: not implemented"; return *new(Pointer) }

// SetPosition sets the position of the value.
func (l *Locator) SetPosition(loc Position) { _ = "STUB: not implemented"; return }

// Position returns the position of the value if it is set.
func (l Locator) Position() (Position, bool) {
	_ = "STUB: not implemented"
	return *

	// Key tries to find the child node using given key and returns its position.
	//
	// See Key method of Position.
	new(Position), false
}

func (l Locator) Key(key string) (loc Locator) { _ = "STUB: not implemented"; return *new(Locator) }

// Field tries to find the child node using given key and returns its position.
//
// See Field method of Position.
func (l Locator) Field(key string) (loc Locator) { _ = "STUB: not implemented"; return *new(Locator) }

// Index tries to find the child node using given index and returns its position.
//
// See Index method of Position.
func (l Locator) Index(idx int) (loc Locator) { _ = "STUB: not implemented"; return *new(Locator) }

// UnmarshalYAML implements yaml.Unmarshaler.
func (l *Locator) UnmarshalYAML(n *yaml.Node) error { _ = "STUB: not implemented"; return nil }
