// Package xslices provides some generic utilities missed from x/exp/slices.
package xslices

// Filter performs in-place filtering of a slice.
func Filter[S ~[]E, E any](s S, keep func(E) bool) S { _ = "STUB: not implemented"; return *new(S) }

// FindFunc returns the first element satisfying the predicate.
func FindFunc[S ~[]E, E any](s S, equal func(E) bool) (r E, _ bool) {
	_ = "STUB: not implemented"
	return *new(E), false
}
