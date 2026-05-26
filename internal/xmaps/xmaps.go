// Package xmaps provides some generic utilities missed from x/exp/maps.
package xmaps

import (
	"golang.org/x/exp/constraints"
)

// SortedKeys returns a sorted slice of keys in the map.
func SortedKeys[M ~map[K]V, K constraints.Ordered, V any](m M) []K {
	_ = "STUB: not implemented"
	return nil
}
