package location

import (
	"github.com/go-faster/yaml"
)

// Position is a value position.
type Position struct {
	Line, Column int
	Node         *yaml.Node
}

// FromNode sets the position of the value from the given node.
func (p *Position) FromNode(node *yaml.Node) { _ = "STUB: not implemented"; return }

func (p Position) mapping() ([]*yaml.Node, bool) { _ = "STUB: not implemented"; return nil, false }

// Key tries to find the child node using given key and returns its position.
// If such node is not found or parent node is not a mapping, Key returns position of the parent node.
//
// NOTE: child position will point to the key node, not to the value node.
// Use Field if you want position of the value.
func (p Position) Key(key string) (loc Position) { _ = "STUB: not implemented"; return *new(Position) }

// Field tries to find the child node using given key and returns its position.
// If such node is not found or parent node is not a mapping, Field returns position of the parent node.
//
// NOTE: child position will point to the value node, not to the key node.
// Use Key if you want position of the key.
func (p Position) Field(key string) (loc Position) {
	_ = "STUB: not implemented"
	return *new(Position)
}

// Index tries to find the child node using given index and returns its position.
// If such node is not found or parent node is not a sequence, Field returns position of the parent node.
func (p Position) Index(idx int) (loc Position) { _ = "STUB: not implemented"; return *new(Position) }

// String implements fmt.Stringer.
func (p Position) String() string { _ = "STUB: not implemented"; return "" }

// WithFilename prints the position with the given filename.
//
// If filename is empty, the position is printed as is.
func (p Position) WithFilename(filename string) string { _ = "STUB: not implemented"; return "" }
