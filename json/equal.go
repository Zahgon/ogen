package json

import (
	"github.com/go-faster/jx"
)

type compare struct {
	left, right *jx.Decoder
}

func (c compare) equalBool() (bool, error) { _ = "STUB: not implemented"; return false, nil }

func (c compare) equalString() (bool, error) { _ = "STUB: not implemented"; return false, nil }

func (c compare) equalNumber() (bool, error) { _ = "STUB: not implemented"; return false, nil }

// Fast path comparing.

// Compare like byte slices.

// If both values are integer, non-zero and are not equal as byte slice,
// so they are not equal as numbers.

func (c compare) equalArray() (bool, error) { _ = "STUB: not implemented"; return false, nil }

// Left array is bigger than right.

// Right array is bigger than left.

func (c compare) equalObject() (bool, error) {
	_ = "STUB: not implemented"
	// TODO(tdakkota): is there a more efficient way?
	return false, nil
}

// Right object is smaller than left.

func (c compare) equal() (bool, error) { _ = "STUB: not implemented"; return false, nil }

// lt is equal to rt, so values are equal.

// Equal compares two JSON values.
func Equal(a, b []byte) (bool, error) { _ = "STUB: not implemented"; return false, nil }
