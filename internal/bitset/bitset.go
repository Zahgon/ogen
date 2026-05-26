// Package bitset implements a byte-slice based bitset
package bitset

// Bitset is a byte-slice based bitset.
type Bitset []uint8

// Set sets the bit by given index.
func (r *Bitset) Set(i int, v bool) { _ = "STUB: not implemented"; return }

// Build builds a bitset from slice using given predicate.
func Build[T any](s []T, cb func(int, T) bool) (r Bitset) {
	_ = "STUB: not implemented"
	return *new(Bitset)
}
